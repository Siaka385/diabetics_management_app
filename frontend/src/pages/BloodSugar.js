           /* BLOOD SUGAR MONITOR*/
         // Blood Sugar Tracking Logic
window.onload=()=>{
    const bloodSugarForm = document.getElementById('bloodSugarForm');
    const alertMessage = document.getElementById('alertMessage');
    const recentLogsList = document.getElementById('recentLogsList');

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

    // Blood sugar level recommendations
    function getBloodSugarRecommendation(level, time) {
        const recommendations = {
            fasting: { low: 80, high: 130 },
            before_meal: { low: 80, high: 130 },
            after_meal: { low: 80, high: 180 },
            bedtime: { low: 100, high: 140 }
        };

        const range = recommendations[time];
        if (level > range.high) {
            return {
                type: 'high',
                message: `Your blood sugar is high for ${time}. Consider: 
                - Drink water
                - Take a short walk
                - Check insulin/medication
                - Consult your healthcare provider`
            };
        } else if (level < range.low) {
            return {
                type: 'low',
                message: `Your blood sugar is low for ${time}. Consider:
                - Have a small snack
                - Drink juice or glucose drink
                - Rest and monitor closely`
            };
        }
        return { type: 'normal', message: 'Your blood sugar is within the target range.' };
    }

    // Form submission handler
    bloodSugarForm.addEventListener('submit', function(e) {
        e.preventDefault();
        const bloodSugarLevel = parseInt(document.getElementById('bloodSugarLevel').value);
        const measurementTime = document.getElementById('measurementTime').value;

        const recommendation = getBloodSugarRecommendation(bloodSugarLevel, measurementTime);
        
        alertMessage.textContent = recommendation.message;
        alertMessage.className = `alert alert-${recommendation.type}`;

        // Add new log
        bloodSugarLogs.unshift({ 
            value: bloodSugarLevel, 
            time: measurementTime, 
            date: new Date().toISOString().split('T')[0] 
        });

        renderRecentLogs();
        updateBloodSugarChart();
    });

    // Initialize chart
    function updateBloodSugarChart() {
        const chartDom = document.getElementById('bloodSugarChart');
        const myChart = echarts.init(chartDom);
        
        const option = {
            tooltip: { trigger: 'axis' },
            xAxis: {
                type: 'category',
                data: bloodSugarLogs.slice(0, 7).map(log => log.date).reverse()
            },
            yAxis: {
                type: 'value',
                min: 50,
                max: 250
            },
            series: [{
                data: bloodSugarLogs.slice(0, 7).map(log => log.value).reverse(),
                type: 'line',
                smooth: true,
                itemStyle: { color: '#1e3a8a' }
            }]
        };

        myChart.setOption(option);
    }
    
    // Initial render
    renderRecentLogs();
    updateBloodSugarChart();


}