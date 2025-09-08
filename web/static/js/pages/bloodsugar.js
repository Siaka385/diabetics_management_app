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
                            <div id="simpleChart" class="w-full h-96"></div>
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
                </div>
            </div>
        `;
        
        // Initialize simple chart
        initializeSimpleChart();
        
        // Add form handler
        const form = document.getElementById('bloodSugarForm');
        if (form) {
            form.addEventListener('submit', handleBloodSugarSubmit);
        }
        
    } catch (error) {
        console.error('Error rendering blood sugar page:', error);
        window.location.href = '/login';
    }
}

// Simple chart data storage
let chartData = [
    { label: 'Mon', value: 120 },
    { label: 'Tue', value: 135 },
    { label: 'Wed', value: 128 },
    { label: 'Thu', value: 142 },
    { label: 'Fri', value: 118 },
    { label: 'Sat', value: 125 },
    { label: 'Sun', value: 130 }
];

function initializeSimpleChart() {
    const container = document.getElementById('simpleChart');
    if (!container) return;
    
    renderSimpleChart(container);
}

function renderSimpleChart(container) {
    const maxValue = Math.max(...chartData.map(d => d.value));
    const minValue = Math.min(...chartData.map(d => d.value));
    const range = maxValue - minValue || 1;
    const chartHeight = 300;
    
    let html = `
        <div class="relative bg-gray-50 rounded-lg p-4" style="height: ${chartHeight + 80}px;">
            <!-- Y-axis labels -->
            <div class="absolute left-0 top-4 bottom-16 w-12 flex flex-col justify-between text-xs text-gray-500">
                <span>${Math.round(maxValue + 20)}</span>
                <span>${Math.round(maxValue)}</span>
                <span>${Math.round((maxValue + minValue) / 2)}</span>
                <span>${Math.round(minValue)}</span>
                <span>${Math.round(minValue - 20)}</span>
            </div>
            
            <!-- Chart area -->
            <div class="ml-16 mr-4 relative" style="height: ${chartHeight}px;">
                <!-- Grid lines -->
                <div class="absolute inset-0">
                    <div class="h-full flex flex-col justify-between">
                        <div class="h-px bg-gray-200"></div>
                        <div class="h-px bg-gray-200"></div>
                        <div class="h-px bg-gray-200"></div>
                        <div class="h-px bg-gray-200"></div>
                        <div class="h-px bg-gray-200"></div>
                    </div>
                </div>
                
                <!-- Data points and line -->
                <svg class="absolute inset-0 w-full h-full" viewBox="0 0 400 ${chartHeight}">
                    <defs>
                        <linearGradient id="gradient" x1="0%" y1="0%" x2="0%" y2="100%">
                            <stop offset="0%" style="stop-color:#3B82F6;stop-opacity:0.3" />
                            <stop offset="100%" style="stop-color:#3B82F6;stop-opacity:0" />
                        </linearGradient>
                    </defs>
                    ${generateSVGPath(chartData, chartHeight, range, minValue)}
                </svg>
                
                <!-- Data points -->
                <div class="absolute inset-0 flex items-end justify-between">
                    ${chartData.map(point => {
                        const height = ((point.value - minValue) / range) * (chartHeight - 40) + 20;
                        return `
                            <div class="flex flex-col items-center">
                                <div class="w-3 h-3 bg-blue-600 rounded-full relative" style="margin-bottom: ${chartHeight - height - 6}px;">
                                    <div class="absolute -top-8 left-1/2 transform -translate-x-1/2 bg-gray-800 text-white text-xs px-2 py-1 rounded opacity-0 hover:opacity-100 transition-opacity">
                                        ${point.value} mg/dL
                                    </div>
                                </div>
                            </div>
                        `;
                    }).join('')}
                </div>
            </div>
            
            <!-- X-axis labels -->
            <div class="ml-16 mr-4 mt-2 flex justify-between text-xs text-gray-500">
                ${chartData.map(point => `<span>${point.label}</span>`).join('')}
            </div>
        </div>
    `;
    
    container.innerHTML = html;
}

function generateSVGPath(data, height, range, minValue) {
    if (data.length < 2) return '';
    
    const svgWidth = 400; // Fixed width in pixels
    const svgHeight = height;
    const stepWidth = svgWidth / (data.length - 1);
    
    let pathData = '';
    let areaData = '';
    
    data.forEach((point, index) => {
        const x = index * stepWidth;
        const y = svgHeight - (((point.value - minValue) / range) * (svgHeight * 0.8) + (svgHeight * 0.1));
        
        if (index === 0) {
            pathData += `M ${x} ${y}`;
            areaData += `M ${x} ${y}`;
        } else {
            pathData += ` L ${x} ${y}`;
            areaData += ` L ${x} ${y}`;
        }
    });
    
    // Close the area path
    areaData += ` L ${svgWidth} ${svgHeight} L 0 ${svgHeight} Z`;
    
    return `
        <path d="${areaData}" fill="url(#gradient)" />
        <path d="${pathData}" stroke="#3B82F6" stroke-width="3" fill="none" />
    `;
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
    
    chartData.push({ label: timeLabel, value: level });
    
    // Keep only the last 10 data points
    if (chartData.length > 10) {
        chartData.shift();
    }
    
    // Re-render the simple chart
    const container = document.getElementById('simpleChart');
    if (container) {
        renderSimpleChart(container);
    }
    
    // Add to recent logs
    addToRecentLogs(level, time);
    
    // Reset form
    e.target.reset();
    
    // Show success message
    showAlert('Blood sugar reading logged successfully!', 'success');
}

function addToRecentLogs(level, time) {
    const logsList = document.getElementById('recentLogsList');
    if (!logsList) return;
    
    const timeFormatted = time.replace('_', ' ').replace(/\b\w/g, l => l.toUpperCase());
    const newLog = document.createElement('div');
    newLog.className = 'p-3 bg-gray-50 rounded-lg';
    newLog.innerHTML = `
        <div class="flex justify-between">
            <span class="font-medium">${level} mg/dL</span>
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

function showAlert(message, type) {
    const alert = document.getElementById('alertMessage');
    if (!alert) return;
    
    alert.classList.remove('hidden');
    alert.className = `mt-4 p-3 rounded-lg ${
        type === 'success' 
            ? 'bg-green-100 text-green-700 border border-green-200' 
            : 'bg-red-100 text-red-700 border border-red-200'
    }`;
    alert.textContent = message;
    
    setTimeout(() => {
        alert.classList.add('hidden');
    }, 3000);
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