document.addEventListener('DOMContentLoaded', () => {
    document.getElementById('mealLogForm').addEventListener('submit', async function (event) {
        event.preventDefault();

        // Get form values
        const mealType = document.getElementById('mealType').value;
        const foodItem = document.getElementById('foodItem').value;
        const weight = parseFloat(document.getElementById('weight').value);
        const proportion = parseFloat(document.getElementById('proportion').value);

        // Create the JSON data to send
        const mealData = {
            MealType: mealType,
            FoodItem: foodItem,
            Weight: weight,
            Proportion: proportion,
        };

        // Send AJAX POST request to the server
        await fetch('/nutrition/logmeal', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(mealData),
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                console.log('Received JSON from server:', data);
                // Update the table with the new meal information
                if (data && data.DietProfile) {
                    addMealToTable(data.DietProfile);
                } else {
                    console.log('Did not find diet profile in server response')
                }
            })
            .catch(error => {
                console.error('Error:', error);
                alert('Failed to log meal. Check the console for errors.'); // Notify user of error
            });
    });

    function addMealToTable(dietProfile) {
        const mealList = document.getElementById('mealList');

        // Create a new table row
        const newRow = document.createElement('tr');

        // Add data cells for Meal Type, Food Item, Calories, and Carbs
        const mealTypeCell = document.createElement('td');
        mealTypeCell.textContent = dietProfile.MealType;
        newRow.appendChild(mealTypeCell);

        const foodItemCell = document.createElement('td');
        foodItemCell.textContent = dietProfile.FoodName;
        newRow.appendChild(foodItemCell);

        const caloriesCell = document.createElement('td');
        caloriesCell.textContent = dietProfile.CaloriesIntake;
        newRow.appendChild(caloriesCell);

        const carbsCell = document.createElement('td');
        carbsCell.textContent = dietProfile.CarbIntake;
        newRow.appendChild(carbsCell);

        // Add new row to the table
        mealList.appendChild(newRow);
    };

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

    // Meal Plan Generation Logic
    function generateMealPlan(e) {
        e.preventDefault();

        const mp_duration = document.querySelector('#mealPlanDurationTags .selected')?.dataset.value;
        const mp_type = document.querySelector('#mealTypeTags .selected')?.dataset.value;
        const mp_diet_pref = document.querySelector('#dietPreferenceTags .selected')?.dataset.value;
        const mp_preferred_foods = document.getElementById('preferredFoods').value.split(',').map(f => f.trim()).filter(f => f);
        const mp_avoid_foods = document.getElementById('avoidFoods').value.split(',').map(f => f.trim()).filter(f => f);

        const mealPlanBody = document.getElementById('mealPlanBody');
        mealPlanBody.innerHTML = '';
        const genInfo = {
            mp_duration,
            mp_type,
            mp_diet_pref,
            mp_preferred_foods,
            mp_avoid_foods,
        }
        // Meal Plan Generation Logic with Diabetes-Friendly Considerations
        fetch("/nutrition/mealplan", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(genInfo),
        });

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
        const selectedPlan = mealPlans[mp_diet_pref] || mealPlans['low-carb'];

        // Filter meals based on selection
        let mealsToDisplay = selectedPlan;
        if (mp_type && mp_type !== 'all-meals') {
            mealsToDisplay = selectedPlan.filter(m => m.meal.toLowerCase() === mp_type);
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