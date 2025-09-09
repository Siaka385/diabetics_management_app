export async function renderBloodsugar() {
    try {
        const response = await fetch('/bloodsugar');
        if (!response.ok) throw new Error('Failed to fetch');
        
        const userData = await response.json();
        const app = document.getElementById('app');
        if (!app) return;
        
        app.innerHTML = `
            <div class="min-h-screen bg-gray-50 flex">
                ${renderSidebar(userData, '/bloodsugar')}
                <div class="flex-1 p-8">
                    <h1 class="text-3xl font-bold mb-8">Blood Sugar Tracking</h1>
                    
                    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
                        <div class="bg-white p-6 rounded-xl shadow-sm border">
                            <h2 class="text-xl font-semibold mb-6">Blood Sugar Trends</h2>
                            <div id="chartContainer" class="w-full h-96"></div>
                        </div>
                        
                        <div class="bg-white p-6 rounded-xl shadow-sm border">
                            <h2 class="text-xl font-semibold mb-6">Log Blood Sugar</h2>
                            <form id="bloodSugarForm" class="space-y-4">
                                <div>
                                    <label for="bloodSugarLevel" class="block mb-2 font-medium text-gray-700">Blood Sugar Level (mg/dL)</label>
                                    <input type="number" id="bloodSugarLevel" required min="50" max="500"
                                        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:border-blue-500">
                                </div>
                                
                                <div>
                                    <label for="measurementTime" class="block mb-2 font-medium text-gray-700">Time of Measurement</label>
                                    <select id="measurementTime" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:border-blue-500">
                                        <option value="fasting">Fasting</option>
                                        <option value="before_meal">Before Meal</option>
                                        <option value="after_meal">After Meal</option>
                                        <option value="bedtime">Bedtime</option>
                                    </select>
                                </div>
                                
                                <button type="submit" class="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 transition-colors">
                                    Log Reading
                                </button>
                            </form>
                            
                            <div id="alertMessage" class="mt-4 hidden"></div>
                            
                            <div class="mt-6">
                                <h3 class="text-lg font-semibold mb-4">Recent Logs</h3>
                                <div id="recentLogsList" class="space-y-2">
                                    <div class="p-3 bg-gray-50 rounded-lg">
                                        <div class="flex justify-between">
                                            <span class="font-medium">128 mg/dL</span>
                                            <span class="text-sm text-gray-500">Fasting - Today 8:00 AM</span>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    
                    <!-- Blood Sugar Interpretation Summary -->
                    <div class="mt-8 bg-white p-6 rounded-xl shadow-sm border">
                        <h2 class="text-xl font-semibold mb-6">Understanding Your Blood Sugar Readings</h2>
                        
                        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                            <div>
                                <h3 class="text-lg font-medium mb-4 text-blue-600">Normal Ranges (mg/dL)</h3>
                                <div class="space-y-3">
                                    <div class="flex justify-between p-3 bg-green-50 rounded-lg border-l-4 border-green-400">
                                        <span class="font-medium">Fasting:</span>
                                        <span>70-100</span>
                                    </div>
                                    <div class="flex justify-between p-3 bg-green-50 rounded-lg border-l-4 border-green-400">
                                        <span class="font-medium">Before Meal:</span>
                                        <span>70-130</span>
                                    </div>
                                    <div class="flex justify-between p-3 bg-green-50 rounded-lg border-l-4 border-green-400">
                                        <span class="font-medium">2hrs After Meal:</span>
                                        <span>&lt; 180</span>
                                    </div>
                                    <div class="flex justify-between p-3 bg-green-50 rounded-lg border-l-4 border-green-400">
                                        <span class="font-medium">Bedtime:</span>
                                        <span>100-140</span>
                                    </div>
                                </div>
                            </div>
                            
                            <div>
                                <h3 class="text-lg font-medium mb-4 text-amber-600">Warning Signs</h3>
                                <div class="space-y-3">
                                    <div class="p-3 bg-yellow-50 rounded-lg border-l-4 border-yellow-400">
                                        <div class="font-medium text-yellow-800">Pre-diabetes:</div>
                                        <div class="text-sm text-yellow-700">Fasting: 100-125 mg/dL</div>
                                    </div>
                                    <div class="p-3 bg-red-50 rounded-lg border-l-4 border-red-400">
                                        <div class="font-medium text-red-800">Diabetes:</div>
                                        <div class="text-sm text-red-700">Fasting: ≥126 mg/dL</div>
                                    </div>
                                    <div class="p-3 bg-red-50 rounded-lg border-l-4 border-red-400">
                                        <div class="font-medium text-red-800">Hypoglycemia:</div>
                                        <div class="text-sm text-red-700">&lt; 70 mg/dL</div>
                                    </div>
                                    <div class="p-3 bg-red-50 rounded-lg border-l-4 border-red-400">
                                        <div class="font-medium text-red-800">Severe High:</div>
                                        <div class="text-sm text-red-700">&gt; 250 mg/dL</div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        
                        <div class="mt-6 p-4 bg-blue-50 rounded-lg border-l-4 border-blue-400">
                            <h4 class="font-medium text-blue-800 mb-2">Important Notes:</h4>
                            <ul class="text-sm text-blue-700 space-y-1">
                                <li>• These are general guidelines - your target ranges may be different based on your condition</li>
                                <li>• Consistently high or low readings should be discussed with your healthcare provider</li>
                                <li>• Factors like stress, illness, medications, and food can affect readings</li>
                                <li>• Track patterns and trends, not just individual readings</li>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
        `;
        
        initializeChart();
        
        const form = document.getElementById('bloodSugarForm');
        if (form) form.addEventListener('submit', handleBloodSugarSubmit);
        
    } catch (error) {
        console.error('Error rendering blood sugar page:', error);
        window.location.href = '/login';
    }
}

// Chart data storage
let chartData = [
    { label: 'Mon', value: 120, time: 'fasting', timestamp: Date.now() - 6*24*60*60*1000 },
    { label: 'Tue', value: 135, time: 'after_meal', timestamp: Date.now() - 5*24*60*60*1000 },
    { label: 'Wed', value: 128, time: 'fasting', timestamp: Date.now() - 4*24*60*60*1000 },
    { label: 'Thu', value: 142, time: 'after_meal', timestamp: Date.now() - 3*24*60*60*1000 },
    { label: 'Fri', value: 118, time: 'fasting', timestamp: Date.now() - 2*24*60*60*1000 },
    { label: 'Sat', value: 125, time: 'before_meal', timestamp: Date.now() - 1*24*60*60*1000 },
    { label: 'Sun', value: 130, time: 'fasting', timestamp: Date.now() }
];

function initializeChart() {
    const container = document.getElementById('chartContainer');
    if (!container) return;
    renderChart(container);
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

function handleBloodSugarSubmit(e) {
    e.preventDefault();
    
    const levelInput = document.getElementById('bloodSugarLevel');
    const timeInput = document.getElementById('measurementTime');
    
    if (!levelInput || !timeInput) return;
    
    const level = parseInt(levelInput.value);
    const time = timeInput.value;
    
    if (isNaN(level) || level < 50 || level > 500) {
        showAlert('Please enter a valid blood sugar level between 50-500 mg/dL', 'error');
        return;
    }
    
    // Add new data point
    const now = new Date();
    const timeLabel = now.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });

    chartData.push({ label: timeLabel, value: level, time: time, timestamp: Date.now() });
    
    // Keep only the last 10 data points
    if (chartData.length > 10) {
        chartData.shift();
    }
    
    // Re-render the chart
    const container = document.getElementById('chartContainer');
    if (container) {
        renderChart(container);
    }
    
    // Add to recent logs
    addToRecentLogs(level, time);
    
    // Reset form
    e.target.reset();
    
    // Show success message with interpretation
    const interpretation = interpretReading(level, time);
    showAlert(`Blood sugar reading logged: ${level} mg/dL - ${interpretation}`, 
             interpretation.includes('normal') ? 'success' : 'warning');
}

function interpretReading(value, timeType) {
    if (value < 70) return 'Low (Hypoglycemia) - Consider treating immediately';
    if (value > 250) return 'Very High - Consult healthcare provider';
    
    switch (timeType) {
        case 'fasting':
            if (value <= 100) return 'Normal fasting level';
            if (value <= 125) return 'Pre-diabetic range';
            return 'Diabetic range';
        
        case 'before_meal':
            if (value <= 130) return 'Normal pre-meal level';
            return 'Above target range';
        
        case 'after_meal':
            if (value < 180) return 'Normal post-meal level';
            return 'Above target range';
        
        case 'bedtime':
            if (value >= 100 && value <= 140) return 'Normal bedtime level';
            if (value < 100) return 'Slightly low for bedtime';
            return 'Above target range';
        
        default:
            return 'Reading logged';
    }
}

function addToRecentLogs(level, time) {
    const logsList = document.getElementById('recentLogsList');
    if (!logsList) return;
    
    const timeFormatted = formatTimeType(time);
    const interpretation = interpretReading(level, time);
    const color = getLogColor(level, time);
    
    const newLog = document.createElement('div');
    newLog.className = `p-3 rounded-lg ${color}`;
    newLog.innerHTML = `
        <div class="flex justify-between items-start">
            <div>
                <span class="font-medium">${level} mg/dL</span>
                <div class="text-xs text-gray-600 mt-1">${interpretation}</div>
            </div>
            <span class="text-sm text-gray-500">${timeFormatted} - Just now</span>
        </div>
    `;
    
    logsList.insertBefore(newLog, logsList.firstChild);
    
    // Keep only the last 5 recent logs visible
    const logs = logsList.children;
    while (logs.length > 5) {
        logsList.removeChild(logs[logs.length - 1]);
    }
}

function getLogColor(value, timeType) {
    const pointColor = getPointColor(value, timeType);
    
    if (pointColor.includes('green')) return 'bg-green-50 border-l-4 border-green-400';
    if (pointColor.includes('yellow')) return 'bg-yellow-50 border-l-4 border-yellow-400';
    if (pointColor.includes('red')) return 'bg-red-50 border-l-4 border-red-400';
    
    return 'bg-gray-50';
}

function showAlert(message, type) {
    const alert = document.getElementById('alertMessage');
    if (!alert) return;
    
    alert.classList.remove('hidden');
    
    let colorClass = '';
    switch (type) {
        case 'success':
            colorClass = 'bg-green-100 text-green-700 border border-green-200';
            break;
        case 'warning':
            colorClass = 'bg-yellow-100 text-yellow-700 border border-yellow-200';
            break;
        case 'error':
            colorClass = 'bg-red-100 text-red-700 border border-red-200';
            break;
        default:
            colorClass = 'bg-blue-100 text-blue-700 border border-blue-200';
    }
    
    alert.className = `mt-4 p-3 rounded-lg ${colorClass}`;
    alert.textContent = message;
    
    setTimeout(() => {
        alert.classList.add('hidden');
    }, 5000);
}

function renderSidebar(userData, currentPage) {
    return `
        <div class="w-64 bg-white shadow-lg">
            <div class="p-6 border-b">
                <div class="text-center">
                    <div class="w-16 h-16 bg-blue-600 rounded-full flex items-center justify-center text-white text-2xl font-bold mx-auto mb-3">
                        ${userData.Abbrev || userData.name?.charAt(0) || 'U'}
                    </div>
                    <h2 class="text-lg font-semibold">${userData.Name || userData.name || 'User'}</h2>
                </div>
            </div>
            <nav class="p-4">
                <ul class="space-y-2">
                    <li><a href="/dashboard" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg ${currentPage === '/dashboard' ? 'text-blue-600 bg-blue-50' : ''}"><span class="mr-3">🏠</span>Dashboard</a></li>
                    <li><a href="/bloodsugar" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg ${currentPage === '/bloodsugar' ? 'text-blue-600 bg-blue-50' : ''}"><span class="mr-3">📊</span>Blood Sugar</a></li>
                    <li><a href="/nutrition" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg ${currentPage === '/nutrition' ? 'text-blue-600 bg-blue-50' : ''}"><span class="mr-3">🥗</span>Nutrition</a></li>
                    <li><a href="/medication" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg ${currentPage === '/medication' ? 'text-blue-600 bg-blue-50' : ''}"><span class="mr-3">💊</span>Medications</a></li>
                    <li><a href="/blog" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg ${currentPage === '/blog' ? 'text-blue-600 bg-blue-50' : ''}"><span class="mr-3">📚</span>Blog</a></li>
                    <li><a href="/support" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg ${currentPage === '/support' ? 'text-blue-600 bg-blue-50' : ''}"><span class="mr-3">🤝</span>Support</a></li>
                    <li><a href="/education" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg ${currentPage === '/education' ? 'text-blue-600 bg-blue-50' : ''}"><span class="mr-3">📖</span>Education</a></li>
                    <li><a href="/glucose-tracker" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg ${currentPage === '/glucose-tracker' ? 'text-blue-600 bg-blue-50' : ''}"><span class="mr-3">📈</span>Glucose Tracker</a></li>
                </ul>
            </nav>
            <div class="p-4 border-t mt-auto">
                <button onclick="fetch('/auth/signout', {method: 'POST'}).then(() => window.location.href = '/')" 
                    class="w-full flex items-center px-4 py-3 text-red-600 hover:bg-red-50 rounded-lg">
                    <span class="mr-3">🚪</span>Logout
                </button>
            </div>
        </div>
    `;
}

function getPointColor(value, timeType) {
    if (value < 70) return 'bg-red-500'; // Hypoglycemia
    if (value > 250) return 'bg-red-600'; // Dangerously high
    
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

function renderChart(container) {
    // Sort data by timestamp
    chartData.sort((a, b) => a.timestamp - b.timestamp);

    // Calculate time range
    const timestamps = chartData.map(d => d.timestamp);
    const minTimestamp = Math.min(...timestamps);
    const maxTimestamp = Math.max(...timestamps);
    const timeRange = maxTimestamp - minTimestamp;

    const maxValue = Math.max(...chartData.map(d => d.value)) + 20;
    const minValue = Math.max(0, Math.min(...chartData.map(d => d.value)) - 20);
    const range = maxValue - minValue;
    const chartHeight = 300;
    const padding = { top: 20, right: 20, bottom: 40, left: 60 };
    
    // Calculate Y-axis scale ticks
    const tickCount = 6;
    const tickStep = range / (tickCount - 1);
    const yTicks = [];
    for (let i = 0; i < tickCount; i++) {
        yTicks.push(Math.round(minValue + (i * tickStep)));
    }
    
    container.innerHTML = `
        <div class="relative bg-gray-50 rounded-lg p-4" style="height:${chartHeight + padding.top + padding.bottom}px;">
            <!-- Y-axis scale -->
            <div class="absolute left-0 top-0" style="width:${padding.left}px; height:${chartHeight + padding.top + padding.bottom}px;">
                ${yTicks.map((tick, index) => {
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
                }).join('')}
                
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
                    
                    <!-- Target range highlight -->
                    <linearGradient id="targetGradient" x1="0%" y1="0%" x2="0%" y2="100%">
                        <stop offset="0%" style="stop-color:#10B981;stop-opacity:0.1" />
                        <stop offset="100%" style="stop-color:#10B981;stop-opacity:0.05" />
                    </linearGradient>
                </defs>
                
                <!-- Target range background -->
                ${generateTargetRangeBackground(chartHeight, range, minValue)}
                
                <!-- Chart path and area -->
                ${generateSVGPath(chartData, chartHeight, range, minValue, minTimestamp, maxTimestamp, timeRange)}
            </svg>
            
            <!-- Data points -->
            ${chartData.map((point, index) => {
                const x = timeRange > 0 ? ((point.timestamp - minTimestamp) / timeRange) * (container.clientWidth - padding.left - padding.right) : (index / (chartData.length - 1)) * (container.clientWidth - padding.left - padding.right);
                const y = chartHeight - ((point.value - minValue) / range) * chartHeight;
                const color = getPointColor(point.value, point.time);
                const dateStr = new Date(point.timestamp).toLocaleDateString();
                return `
                    <div class="absolute w-4 h-4 ${color} rounded-full border-2 border-white shadow-lg cursor-pointer hover:scale-110 transition-transform"
                         style="left:${padding.left + x}px; top:${padding.top + y - 8}px; transform:translateX(-50%);"
                         title="${point.value} mg/dL - ${formatTimeType(point.time)} - ${dateStr} ${point.label}"></div>
                `;
            }).join('')}
            
            <!-- X-axis labels -->
            <div class="flex justify-between text-xs text-gray-500 mt-2"
                 style="margin-left:${padding.left}px; margin-right:${padding.right}px;">
                ${chartData.map(p => {
                    const date = new Date(p.timestamp);
                    const dateStr = date.toLocaleDateString([], { month: 'short', day: 'numeric' });
                    return `<span class="text-center">${dateStr}</span>`;
                }).join('')}
            </div>
            
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

function generateTargetRangeBackground(height, range, minValue) {
    // Highlight the general target range (70-180 mg/dL) as a subtle background
    const targetMin = 70;
    const targetMax = 180;
    
    if (targetMax < minValue || targetMin > (minValue + range)) {
        return ''; // Target range not visible in current chart range
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
