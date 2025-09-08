export async function renderNutrition() {
    try {
        const response = await fetch('/nutrition');
        if (!response.ok) throw new Error('Failed to fetch');
        
        const userData = await response.json();
        const app = document.getElementById('app');
        if (!app) return;
        
        app.innerHTML = `
            <div class="min-h-screen bg-gray-50 flex">
                ${renderSidebar(userData, '/nutrition')}
                <div class="flex-1 p-8">
                    <h1 class="text-3xl font-bold mb-8">Advanced Nutrition Tracking</h1>
                    
                    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
                        <div class="bg-white p-6 rounded-xl shadow-sm border">
                            <h2 class="text-xl font-semibold mb-6">Log Your Daily Meals</h2>
                            <form id="mealLogForm" class="space-y-4">
                                <div>
                                    <label for="mealType" class="block mb-2 font-medium text-gray-700">Meal Type</label>
                                    <select id="mealType" required class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:border-blue-500">
                                        <option value="">Select Meal</option>
                                        <option value="Breakfast">Breakfast</option>
                                        <option value="Lunch">Lunch</option>
                                        <option value="Dinner">Dinner</option>
                                        <option value="Snack">Snack</option>
                                    </select>
                                </div>
                                
                                <div>
                                    <label for="foodItem" class="block mb-2 font-medium text-gray-700">Food Item</label>
                                    <input type="text" id="foodItem" placeholder="Enter food item" required 
                                        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:border-blue-500">
                                </div>
                                
                                <div>
                                    <label for="weight" class="block mb-2 font-medium text-gray-700">Weight (g)</label>
                                    <input type="number" id="weight" placeholder="Weight" required 
                                        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:border-blue-500">
                                </div>
                                
                                <div>
                                    <label for="proportion" class="block mb-2 font-medium text-gray-700">Proportion Of Plate</label>
                                    <input type="number" step="0.1" id="proportion" placeholder="0.5" required 
                                        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:border-blue-500">
                                </div>
                                
                                <button type="submit" class="w-full bg-green-600 text-white py-2 px-4 rounded-md hover:bg-green-700 transition-colors">
                                    Log Meal
                                </button>
                            </form>
                            
                            <div class="mt-6">
                                <h3 class="text-lg font-semibold mb-4">Today's Meals</h3>
                                <div class="overflow-x-auto">
                                    <table class="w-full text-sm">
                                        <thead class="bg-gray-50">
                                            <tr>
                                                <th class="px-3 py-2 text-left">Meal Type</th>
                                                <th class="px-3 py-2 text-left">Food Item</th>
                                                <th class="px-3 py-2 text-left">Calories</th>
                                                <th class="px-3 py-2 text-left">Carbs (g)</th>
                                            </tr>
                                        </thead>
                                        <tbody id="mealList">
                                            <tr class="border-t">
                                                <td class="px-3 py-2">Breakfast</td>
                                                <td class="px-3 py-2">Oatmeal</td>
                                                <td class="px-3 py-2">150</td>
                                                <td class="px-3 py-2">27</td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </div>
                            </div>
                        </div>
                        
                        <div class="bg-white p-6 rounded-xl shadow-sm border">
                            <h2 class="text-xl font-semibold mb-6">Personalized Meal Plan Generator</h2>
                            <form id="mealPlanForm" class="space-y-4">
                                <div>
                                    <label class="block mb-2 font-medium text-gray-700">Meal Plan Duration</label>
                                    <div class="flex flex-wrap gap-2">
                                        <button type="button" class="meal-duration-tag px-3 py-1 bg-gray-200 text-gray-700 rounded-full text-sm hover:bg-blue-100" data-value="single-day">Single Day</button>
                                        <button type="button" class="meal-duration-tag px-3 py-1 bg-gray-200 text-gray-700 rounded-full text-sm hover:bg-blue-100" data-value="whole-week">Whole Week</button>
                                    </div>
                                </div>
                                
                                <div>
                                    <label class="block mb-2 font-medium text-gray-700">Meal Types</label>
                                    <div class="flex flex-wrap gap-2">
                                        <button type="button" class="meal-type-tag px-3 py-1 bg-gray-200 text-gray-700 rounded-full text-sm hover:bg-blue-100" data-value="breakfast">Breakfast</button>
                                        <button type="button" class="meal-type-tag px-3 py-1 bg-gray-200 text-gray-700 rounded-full text-sm hover:bg-blue-100" data-value="lunch">Lunch</button>
                                        <button type="button" class="meal-type-tag px-3 py-1 bg-gray-200 text-gray-700 rounded-full text-sm hover:bg-blue-100" data-value="dinner">Dinner</button>
                                        <button type="button" class="meal-type-tag px-3 py-1 bg-gray-200 text-gray-700 rounded-full text-sm hover:bg-blue-100" data-value="all-meals">All Meals</button>
                                    </div>
                                </div>
                                
                                <div>
                                    <label class="block mb-2 font-medium text-gray-700">Dietary Preferences</label>
                                    <div class="flex flex-wrap gap-2">
                                        <button type="button" class="diet-pref-tag px-3 py-1 bg-gray-200 text-gray-700 rounded-full text-sm hover:bg-blue-100" data-value="low-carb">Low Carb</button>
                                        <button type="button" class="diet-pref-tag px-3 py-1 bg-gray-200 text-gray-700 rounded-full text-sm hover:bg-blue-100" data-value="high-protein">High Protein</button>
                                        <button type="button" class="diet-pref-tag px-3 py-1 bg-gray-200 text-gray-700 rounded-full text-sm hover:bg-blue-100" data-value="vegetarian">Vegetarian</button>
                                        <button type="button" class="diet-pref-tag px-3 py-1 bg-gray-200 text-gray-700 rounded-full text-sm hover:bg-blue-100" data-value="gluten-free">Gluten-Free</button>
                                    </div>
                                </div>
                                
                                <button type="submit" class="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 transition-colors">
                                    Generate Meal Plan
                                </button>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        `;
        
        // Add event handlers
        document.getElementById('mealLogForm').addEventListener('submit', handleMealSubmit);
        document.getElementById('mealPlanForm').addEventListener('submit', handleMealPlanSubmit);
        
        // Add tag selection handlers
        document.querySelectorAll('.meal-duration-tag, .meal-type-tag, .diet-pref-tag').forEach(tag => {
            tag.addEventListener('click', (e) => {
                e.preventDefault();
                tag.classList.toggle('bg-blue-500');
                tag.classList.toggle('text-white');
                tag.classList.toggle('bg-gray-200');
                tag.classList.toggle('text-gray-700');
            });
        });
        
    } catch (error) {
        window.location.href = '/login';
    }
}

function handleMealSubmit(e) {
    e.preventDefault();
    const mealType = document.getElementById('mealType').value;
    const foodItem = document.getElementById('foodItem').value;
    const weight = document.getElementById('weight').value;
    
    // Add to meal list
    const mealList = document.getElementById('mealList');
    const newRow = document.createElement('tr');
    newRow.className = 'border-t';
    newRow.innerHTML = `
        <td class="px-3 py-2">${mealType}</td>
        <td class="px-3 py-2">${foodItem}</td>
        <td class="px-3 py-2">${Math.round(weight * 2)}</td>
        <td class="px-3 py-2">${Math.round(weight * 0.3)}</td>
    `;
    mealList.appendChild(newRow);
    
    e.target.reset();
}

function handleMealPlanSubmit(e) {
    e.preventDefault();
    alert('Meal plan generation feature coming soon!');
}

function renderSidebar(userData, currentPage) {
    return `
        <div class="w-64 bg-white shadow-lg">
            <div class="p-6 border-b">
                <div class="text-center">
                    <div class="w-16 h-16 bg-blue-600 rounded-full flex items-center justify-center text-white text-2xl font-bold mx-auto mb-3">
                        ${userData.Abbrev}
                    </div>
                    <h2 class="text-lg font-semibold">${userData.Name}</h2>
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