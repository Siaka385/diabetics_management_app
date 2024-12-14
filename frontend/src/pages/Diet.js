document.addEventListener('DOMContentLoaded', () => {
    // Preference Tag Selection Logic
    function setupPreferenceTags() {
        const tagSections = ['mealPlanDurationTags', 'mealTypeTags', 'dietPreferenceTags'];
        
        tagSections.forEach(sectionId => {
            const section = document.getElementById(sectionId);
            section.addEventListener('click', (e) => {
                if (e.target.classList.contains('preference-tag')) {
                    // Remove selection from other tags in the same section
                    section.querySelectorAll('.preference-tag').forEach(tag => {
                        tag.classList.remove('selected');
                    });
                    
                    // Select the clicked tag
                    e.target.classList.add('selected');
                }
            });
        });
    }

    // Meal Logging Logic
    function logMeal(e) {
        e.preventDefault();
        
        const mealType = document.getElementById('mealType').value;
        const foodItem = document.getElementById('foodItem').value;
        const calories = document.getElementById('calories').value;
        const carbs = document.getElementById('carbs').value;
        
        const mealList = document.getElementById('mealList');
        const newRow = mealList.insertRow();
        
        newRow.innerHTML = `
            <td>${mealType}</td>
            <td>${foodItem}</td>
            <td>${calories}</td>
            <td>${carbs}</td>
        `;
        
        // Optional: Save to local storage
        saveMealToLocalStorage(mealType, foodItem, calories, carbs);
        
        // Reset form
        e.target.reset();
    }

    // Save Meal to Local Storage
    function saveMealToLocalStorage(mealType, foodItem, calories, carbs) {
        let meals = JSON.parse(localStorage.getItem('meals')) || [];
        meals.push({ mealType, foodItem, calories, carbs });
        localStorage.setItem('meals', JSON.stringify(meals));
    }

    // Meal Plan Generation Logic
    function generateMealPlan(e) {
        e.preventDefault();
        
        const selectedDuration = document.querySelector('#mealPlanDurationTags .selected')?.dataset.value;
        const selectedMealType = document.querySelector('#mealTypeTags .selected')?.dataset.value;
        const selectedDietPref = document.querySelector('#dietPreferenceTags .selected')?.dataset.value;
        const preferredFoods = document.getElementById('preferredFoods').value.split(',').map(f => f.trim()).filter(f => f);
        const avoidFoods = document.getElementById('avoidFoods').value.split(',').map(f => f.trim()).filter(f => f);

        const mealPlanBody = document.getElementById('mealPlanBody');
        mealPlanBody.innerHTML = ''; // Clear previous results
        
        // Meal Plan Generation Logic with Diabetes-Friendly Considerations
        const mealPlans = {
            'low-carb': [
                { meal: 'Breakfast', dish: 'Spinach and Feta Egg White Omelet', localAlternative: 'Tofu Scramble', carbs: 8 },
                { meal: 'Lunch', dish: 'Grilled Chicken Salad', localAlternative: 'Fish Tikka Salad', carbs: 10 },
                { meal: 'Dinner', dish: 'Baked Salmon with Roasted Vegetables', localAlternative: 'Tandoori Fish with Cauliflower Rice', carbs: 12 }
            ],
            'high-protein': [
                { meal: 'Breakfast', dish: 'Greek Yogurt Parfait', localAlternative: 'Paneer Bhurji', carbs: 15 },
                { meal: 'Lunch', dish: 'Turkey and Quinoa Bowl', localAlternative: 'Chicken Kebab with Mint Chutney', carbs: 20 },
                { meal: 'Dinner', dish: 'Lean Beef Stir Fry', localAlternative: 'Mutton Seekh Kebab with Raita', carbs: 18 }
            ],
            'vegetarian': [
                { meal: 'Breakfast', dish: 'Tofu Scramble', localAlternative: 'Masala Dosa', carbs: 25 },
                { meal: 'Lunch', dish: 'Lentil and Vegetable Curry', localAlternative: 'Dal Makhani with Cauliflower Rice', carbs: 30 },
                { meal: 'Dinner', dish: 'Stuffed Bell Peppers', localAlternative: 'Vegetable Biryani with Raita', carbs: 35 }
            ],
            'gluten-free': [
                { meal: 'Breakfast', dish: 'Chia Seed Pudding', localAlternative: 'Sabudana Khichdi', carbs: 20 },
                { meal: 'Lunch', dish: 'Grilled Shrimp with Zucchini Noodles', localAlternative: 'Prawn Curry with Rice', carbs: 25 },
                { meal: 'Dinner', dish: 'Roasted Vegetable Quinoa Bowl', localAlternative: 'Mixed Vegetable Curry', carbs: 22 }
            ]
        };

        // Determine meal plan based on selected preferences
        const selectedPlan = mealPlans[selectedDietPref] || mealPlans['low-carb'];

        // Filter meals based on selection
        let mealsToDisplay = selectedPlan;
        if (selectedMealType && selectedMealType !== 'all-meals') {
            mealsToDisplay = selectedPlan.filter(m => m.meal.toLowerCase() === selectedMealType);
        }

        // Populate meal plan
        mealsToDisplay.forEach(item => {
            const newRow = mealPlanBody.insertRow();
            newRow.innerHTML = `
                <td>${item.meal}</td>
                <td>${item.dish}</td>
                <td>${item.localAlternative}</td>
                <td>${item.carbs}</td>
            `;
        });
    }

    // Initialize Event Listeners
    function initializeEventListeners() {
        setupPreferenceTags();
        
        const mealLogForm = document.getElementById('mealLogForm');
        if (mealLogForm) {
            mealLogForm.addEventListener('submit', logMeal);
        }
        
        const mealPlanForm = document.getElementById('mealPlanPreferencesForm');
        if (mealPlanForm) {
            mealPlanForm.addEventListener('submit', generateMealPlan);
        }
    }

    // Load Saved Meals on Page Load
    function loadSavedMeals() {
        const meals = JSON.parse(localStorage.getItem('meals')) || [];
        const mealList = document.getElementById('mealList');
        
        meals.forEach(meal => {
            const newRow = mealList.insertRow();
            newRow.innerHTML = `
                <td>${meal.mealType}</td>
                <td>${meal.foodItem}</td>
                <td>${meal.calories}</td>
                <td>${meal.carbs}</td>
            `;
        });
    }

    // Initialize the application
    initializeEventListeners();
    loadSavedMeals();
});