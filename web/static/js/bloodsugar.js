// Blood Sugar Tracking Logic
const bloodSugarForm = document.getElementById('bloodSugarForm');
const alertMessage = document.getElementById('alertMessage');
const recentLogsList = document.getElementById('recentLogsList');
const bloodSugarLevelInput = document.getElementById('bloodSugarLevel');
const measurementTimeSelect = document.getElementById('measurementTime');

// Sample blood sugar logs (would typically come from backend/storage)
let bloodSugarLogs = [
    { value: 128, time: 'before_meal', date: '2024-03-15' },
    { value: 115, time: 'fasting', date: '2024-03-14' },
    { value: 142, time: 'after_meal', date: '2024-03-13' }
];

// Render recent logs
function renderRecentLogs() {
    recentLogsList.innerHTML = bloodSugarLogs
        .slice(0, 5)
        .map(log => `
            <div class="log-entry">
                <span>${log.value} mg/dL (${log.time})</span>
                <span>${log.date}</span>
            </div>
        `).join('');
}

// Provide guidance for blood sugar levels
function getBloodSugarGuidance(value) {
    const numValue = Number(value);
    
    // Low Blood Sugar (Hypoglycemia) Guidance
    if (numValue < 70) {
        return {
            type: 'low',
            message: 'Low Blood Sugar: Immediate Action Required!',
            guidance: [
                '1. Consume 15-20 grams of fast-acting carbohydrates immediately:',
                '   - 4 oz (1/2 cup) of fruit juice',
                '   - 4-5 glucose tablets',
                '   - 1 tablespoon of honey or sugar',
                '2. Wait 15 minutes and recheck your blood sugar',
                '3. If still below 70 mg/dL, repeat the process',
                '4. Once blood sugar normalizes, eat a small snack',
                '5. If symptoms persist or you feel unable to treat yourself, seek medical help'
            ]
        };
    } 
    // High Blood Sugar (Hyperglycemia) Guidance
    else if (numValue > 180) {
        return {
            type: 'high',
            message: 'High Blood Sugar: Take Careful Action',
            guidance: [
                '1. Check your ketones if blood sugar is consistently above 240 mg/dL',
                '2. Drink water to help flush out excess sugar',
                '3. Do light exercise if safe and approved by your doctor',
                '4. Take insulin or medication as prescribed by your healthcare provider',
                '5. Avoid high-carb foods',
                '6. Monitor your blood sugar closely',
                '7. If blood sugar remains high (>300 mg/dL) or you have symptoms like:',
                '   - Extreme thirst',
                '   - Frequent urination',
                '   - Nausea and vomiting',
                '   CONTACT YOUR HEALTHCARE PROVIDER IMMEDIATELY'
            ]
        };
    }
    
    // Normal Range
    return {
        type: 'normal',
        message: 'Blood Sugar in Normal Range',
        guidance: [
            'Keep maintaining your current diet and medication routine',
            'Continue regular monitoring and healthy lifestyle practices'
        ]
    };
}

// Validate blood sugar level
function validateBloodSugar(value) {
    const guidance = getBloodSugarGuidance(value);
    
    // Update alert message
    alertMessage.textContent = guidance.message;
    alertMessage.classList.remove('alert-high', 'alert-normal');
    
    if (guidance.type === 'low') {
        alertMessage.classList.add('alert-normal');
    } else if (guidance.type === 'high') {
        alertMessage.classList.add('alert-high');
    }

    // Create guidance list
    const guidanceList = document.createElement('ul');
    guidance.guidance.forEach(step => {
        const li = document.createElement('li');
        li.textContent = step;
        guidanceList.appendChild(li);
    });

    // Clear previous guidance and append new list
    const existingList = alertMessage.querySelector('ul');
    if (existingList) {
        alertMessage.removeChild(existingList);
    }
    alertMessage.appendChild(guidanceList);

    return true;
}

// Rest of the previous script remains the same (addBloodSugarLog, form submission, chart initialization, etc.)

// CSS to style the guidance list
const styleTag = document.createElement('style');
styleTag.textContent = `
    .alert ul {
        margin-top: 10px;
        padding-left: 20px;
        text-align: left;
        font-size: 0.9em;
    }
    .alert ul li {
        margin-bottom: 5px;
        line-height: 1.4;
    }
`;
document.head.appendChild(styleTag);

// Add new blood sugar log
function addBloodSugarLog(value, time) {
    const newLog = {
        value: Number(value),
        time: time,
        date: new Date().toISOString().split('T')[0]
    };
    bloodSugarLogs.unshift(newLog);
    renderRecentLogs();
    updateBloodSugarChart();
}

// Form submission handler
bloodSugarForm.addEventListener('submit', function(e) {
    e.preventDefault();
    const bloodSugarLevel = bloodSugarLevelInput.value;
    const measurementTime = measurementTimeSelect.value;

    if (validateBloodSugar(bloodSugarLevel)) {
        addBloodSugarLog(bloodSugarLevel, measurementTime);
        bloodSugarForm.reset();
    }
});

// Initialize ECharts Blood Sugar Trend Chart
function updateBloodSugarChart() {
    const chartDom = document.getElementById('bloodSugarChart');
    const myChart = echarts.init(chartDom);
    
    const option = {
        tooltip: {
            trigger: 'axis',
            axisPointer: { type: 'line' }
        },
        xAxis: {
            type: 'category',
            data: bloodSugarLogs.slice(0, 7).map(log => log.date).reverse()
        },
        yAxis: {
            type: 'value',
            name: 'Blood Sugar (mg/dL)',
            min: 50,
            max: 250
        },
        series: [{
            name: 'Blood Sugar',
            type: 'line',
            data: bloodSugarLogs.slice(0, 7).map(log => log.value).reverse(),
            itemStyle: {
                color: log => {
                    return log > 180 ? '#ef4444' : (log < 70 ? '#10b981' : '#2563eb');
                }
            },
            lineStyle: {
                color: '#2563eb'
            }
        }]
    };

    myChart.setOption(option);
}

// Initial rendering and chart setup
renderRecentLogs();
updateBloodSugarChart();

// Optional: Add input validation to prevent non-numeric input
bloodSugarLevelInput.addEventListener('input', function(e) {
    e.target.value = e.target.value.replace(/[^0-9]/g, '');
});
