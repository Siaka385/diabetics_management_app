// Graphing Module for Glucose Readings
// Supports Full Timeline, Daily, and Rolling 24h charts
// Provides utilities for filtering, scaling, and interactive features

// Utility functions for data filtering and scaling
export function filterReadingsByDateRange(readings, startDate, endDate) {
    return readings.filter(reading => {
        const readingDate = new Date(reading.timestamp);
        return readingDate >= startDate && readingDate <= endDate;
    });
}

export function filterReadingsByTimeType(readings, timeType) {
    if (!timeType) return readings;
    return readings.filter(reading => reading.time === timeType);
}

export function scaleReadings(readings, targetMin = 0, targetMax = 500) {
    const values = readings.map(r => r.value);
    const min = Math.min(...values);
    const max = Math.max(...values);
    const range = max - min;

    return readings.map(reading => ({
        ...reading,
        scaledValue: range > 0 ? ((reading.value - min) / range) * (targetMax - targetMin) + targetMin : reading.value
    }));
}

// Chart type constants
export const CHART_TYPES = {
    FULL_TIMELINE: 'full_timeline',
    DAILY: 'daily',
    ROLLING_24H: 'rolling_24h'
};

// Main chart rendering function
export function renderGlucoseChart(container, readings, chartType = CHART_TYPES.FULL_TIMELINE) {
    // Sort readings by timestamp
    readings.sort((a, b) => a.timestamp - b.timestamp);

    // Filter based on chart type
    let filteredReadings = readings;
    if (chartType === CHART_TYPES.DAILY) {
        // For daily chart, filter to current day
        const today = new Date();
        const startOfDay = new Date(today.getFullYear(), today.getMonth(), today.getDate());
        const endOfDay = new Date(today.getFullYear(), today.getMonth(), today.getDate(), 23, 59, 59);
        filteredReadings = filterReadingsByDateRange(readings, startOfDay, endOfDay);
    } else if (chartType === CHART_TYPES.ROLLING_24H) {
        // For rolling 24h, filter to last 24 hours
        const now = Date.now();
        const twentyFourHoursAgo = now - 24 * 60 * 60 * 1000;
        filteredReadings = readings.filter(r => r.timestamp >= twentyFourHoursAgo);
    }

    // Calculate dimensions
    const maxValue = Math.max(...filteredReadings.map(d => d.value)) + 20;
    const minValue = Math.max(0, Math.min(...filteredReadings.map(d => d.value)) - 20);
    const range = maxValue - minValue;
    const chartHeight = 300;
    const padding = { top: 20, right: 20, bottom: 40, left: 60 };

    // Calculate time range
    const timestamps = filteredReadings.map(d => d.timestamp);
    const minTimestamp = Math.min(...timestamps);
    const maxTimestamp = Math.max(...timestamps);
    const timeRange = maxTimestamp - minTimestamp;

    container.innerHTML = `
        <div class="relative bg-gray-50 rounded-lg p-4" style="height:${chartHeight + padding.top + padding.bottom}px;">
            <!-- Y-axis scale -->
            <div class="absolute left-0 top-0" style="width:${padding.left}px; height:${chartHeight + padding.top + padding.bottom}px;">
                ${generateYAxisLabels(minValue, maxValue, range, chartHeight, padding)}
                <!-- Y-axis label -->
                <div class="absolute left-2 top-1/2 -rotate-90 text-sm text-gray-600 font-medium whitespace-nowrap"
                     style="transform: translateY(-50%) translateX(-50%) rotate(-90deg); transform-origin: center;">
                    Blood Sugar (mg/dL)
                </div>
            </div>

            <!-- Chart SVG -->
            <svg width="100%" height="${chartHeight}" viewBox="0 0 400 ${chartHeight}" style="margin-left:${padding.left}px; margin-right:${padding.right}px;">
                <defs>
                    <linearGradient id="gradient" x1="0%" y1="0%" x2="0%" y2="100%">
                        <stop offset="0%" style="stop-color:#3B82F6;stop-opacity:0.3" />
                        <stop offset="100%" style="stop-color:#3B82F6;stop-opacity:0" />
                    </linearGradient>
                    <linearGradient id="targetGradient" x1="0%" y1="0%" x2="0%" y2="100%">
                        <stop offset="0%" style="stop-color:#10B981;stop-opacity:0.1" />
                        <stop offset="100%" style="stop-color:#10B981;stop-opacity:0.05" />
                    </linearGradient>
                </defs>

                <!-- Target range background -->
                ${generateTargetRangeBackground(chartHeight, range, minValue)}

                <!-- Chart path and area -->
                ${generateSVGPath(filteredReadings, chartHeight, range, minValue, minTimestamp, maxTimestamp, timeRange)}
            </svg>

            <!-- Data points -->
            ${generateDataPoints(filteredReadings, chartHeight, range, minValue, minTimestamp, maxTimestamp, timeRange, padding)}

            <!-- X-axis labels -->
            ${generateXAxisLabels(filteredReadings, padding, chartType)}

            <!-- Chart legend -->
            <div class="absolute bottom-2 right-4 text-xs text-gray-500">
                <div class="flex items-center space-x-4">
                    <div class="flex items-center">
                        <div class="w-3 h-3 bg-green-500 rounded-full mr-1"></div>
                        <span>Normal</span>
                    </div>
                    <div class="flex items-center">
                        <div class="w-3 h-3 bg-yellow-500 rounded-full mr-1"></div>
                        <span>Caution</span>
                    </div>
                    <div class="flex items-center">
                        <div class="w-3 h-3 bg-red-500 rounded-full mr-1"></div>
                        <span>Alert</span>
                    </div>
                </div>
            </div>
        </div>
    `;
}

function generateYAxisLabels(minValue, maxValue, range, chartHeight, padding) {
    const tickCount = 6;
    const tickStep = range / (tickCount - 1);
    const yTicks = [];
    for (let i = 0; i < tickCount; i++) {
        yTicks.push(Math.round(minValue + (i * tickStep)));
    }

    return yTicks.map((tick, index) => {
        const y = padding.top + chartHeight - ((tick - minValue) / range) * chartHeight;
        const isTargetRange = (tick >= 70 && tick <= 140);
        return `
            <div class="absolute right-2 text-xs ${isTargetRange ? 'text-green-600 font-medium' : 'text-gray-500'}"
                 style="top:${y - 8}px; transform:translateY(-50%);">
                ${tick}
            </div>
            <div class="absolute right-0 w-2 h-px bg-gray-300"
                 style="top:${y}px;"></div>
            <div class="absolute h-px ${isTargetRange ? 'bg-green-200' : 'bg-gray-200'} opacity-50"
                 style="top:${y}px; left:${padding.left}px; right:${padding.right}px; width:calc(100% - ${padding.left + padding.right}px);"></div>
        `;
    }).join('');
}

function generateDataPoints(readings, chartHeight, range, minValue, minTimestamp, maxTimestamp, timeRange, padding) {
    return readings.map((point, index) => {
        const x = timeRange > 0 ? ((point.timestamp - minTimestamp) / timeRange) * (400 - padding.left - padding.right) : (index / (readings.length - 1)) * (400 - padding.left - padding.right);
        const y = chartHeight - ((point.value - minValue) / range) * chartHeight;
        const color = getPointColor(point.value, point.time);
        const dateStr = new Date(point.timestamp).toLocaleDateString();
        const clickHandler = `onclick="window.handlePointClick(event, ${point.timestamp})"`;
        return `
            <div class="absolute w-4 h-4 ${color} rounded-full border-2 border-white shadow-lg cursor-pointer hover:scale-110 transition-transform"
                 style="left:${padding.left + x}px; top:${padding.top + y - 8}px; transform:translateX(-50%);"
                 title="${point.value} mg/dL - ${formatTimeType(point.time)} - ${dateStr} ${point.label}"
                 ${clickHandler}></div>
        `;
    }).join('');
}

function generateXAxisLabels(readings, padding, chartType) {
    if (readings.length === 0) return '';

    // For time-based X-axis, show appropriate time labels
    const timestamps = readings.map(p => p.timestamp);
    const minTimestamp = Math.min(...timestamps);
    const maxTimestamp = Math.max(...timestamps);
    const timeRange = maxTimestamp - minTimestamp;

    // Determine label frequency based on time range
    let labelStep;
    if (timeRange < 24 * 60 * 60 * 1000) { // Less than 24 hours
        labelStep = 4 * 60 * 60 * 1000; // Every 4 hours
    } else if (timeRange < 7 * 24 * 60 * 60 * 1000) { // Less than 7 days
        labelStep = 24 * 60 * 60 * 1000; // Every day
    } else {
        labelStep = 7 * 24 * 60 * 60 * 1000; // Every week
    }

    const labels = [];
    for (let t = minTimestamp; t <= maxTimestamp; t += labelStep) {
        const date = new Date(t);
        let label;
        if (timeRange < 24 * 60 * 60 * 1000) {
            label = date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
        } else if (timeRange < 7 * 24 * 60 * 60 * 1000) {
            label = date.toLocaleDateString([], { month: 'short', day: 'numeric' });
        } else {
            label = date.toLocaleDateString([], { month: 'short', day: 'numeric', year: '2-digit' });
        }
        labels.push(label);
    }

    return `
        <div class="flex justify-between text-xs text-gray-500 mt-2"
             style="margin-left:${padding.left}px; margin-right:${padding.right}px;">
            ${labels.map(label => `<span class="text-center">${label}</span>`).join('')}
        </div>
    `;
}

function generateTargetRangeBackground(height, range, minValue) {
    const targetMin = 70;
    const targetMax = 180;

    if (targetMax < minValue || targetMin > (minValue + range)) {
        return '';
    }

    const clampedMin = Math.max(targetMin, minValue);
    const clampedMax = Math.min(targetMax, minValue + range);

    const yTop = 100 - ((clampedMax - minValue) / range) * 100;
    const yBottom = 100 - ((clampedMin - minValue) / range) * 100;

    return `
        <rect x="0%" y="${yTop}%" width="100%" height="${yBottom - yTop}%"
              fill="url(#targetGradient)" opacity="0.5" />
    `;
}

function generateSVGPath(data, height, range, minValue, minTimestamp, maxTimestamp, timeRange) {
    if (data.length < 2) return '';

    const svgWidth = 400;
    const svgHeight = height;
    let pathData = '';
    let areaData = '';

    data.forEach((point, i) => {
        const x = timeRange > 0 ? ((point.timestamp - minTimestamp) / timeRange) * svgWidth : (i / (data.length - 1)) * svgWidth;
        const y = svgHeight - ((point.value - minValue) / range) * svgHeight;

        if (i === 0) {
            pathData += `M ${x} ${y}`;
            areaData += `M ${x} ${y}`;
        } else {
            pathData += ` L ${x} ${y}`;
            areaData += ` L ${x} ${y}`;
        }
    });

    areaData += ` L ${svgWidth} ${svgHeight} L 0 ${svgHeight} Z`;

    return `
        <path d="${areaData}" fill="url(#gradient)" />
        <path d="${pathData}" stroke="#3B82F6" stroke-width="3" fill="none" stroke-linecap="round" stroke-linejoin="round" />
    `;
}

export function getPointColor(value, timeType) {
    if (value < 70) return 'bg-red-500';
    if (value > 250) return 'bg-red-600';

    switch (timeType) {
        case 'fasting':
            if (value <= 100) return 'bg-green-500';
            if (value <= 125) return 'bg-yellow-500';
            return 'bg-red-500';

        case 'before_meal':
            if (value <= 130) return 'bg-green-500';
            if (value <= 180) return 'bg-yellow-500';
            return 'bg-red-500';

        case 'after_meal':
            if (value < 180) return 'bg-green-500';
            if (value <= 250) return 'bg-yellow-500';
            return 'bg-red-500';

        case 'bedtime':
            if (value >= 100 && value <= 140) return 'bg-green-500';
            if ((value >= 70 && value < 100) || (value > 140 && value <= 180)) return 'bg-yellow-500';
            return 'bg-red-500';

        default:
            return 'bg-blue-500';
    }
}

function formatTimeType(timeType) {
    const timeMap = {
        'fasting': 'Fasting',
        'before_meal': 'Before Meal',
        'after_meal': 'After Meal',
        'bedtime': 'Bedtime'
    };
    return timeMap[timeType] || timeType;
}
