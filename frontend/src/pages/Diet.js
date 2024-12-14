document.addEventListener('DOMContentLoaded', () => {
    const mealLogForm = document.getElementById('mealLogForm');
    
    mealLogForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        
        // Collect form data
        const mealLogData = {
            mealType: document.getElementById('mealType').value,
            foodItem: document.getElementById('foodItem').value,
            weight: parseFloat(document.getElementById('weight').value),
            proportion: parseFloat(document.getElementById('proportion').value)
        };
        
        try {
            // Send data to backend
            const response = await fetch('/nutrition/meal/log', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(mealLogData)
            });
            
            if (!response.ok) {
                throw new Error('Failed to log meal');
            }
            
            const result = await response.json();
            
            console.log("here",JSON.stringify(result));
            // Handle successful response
            displayMealLogSuccess(result);
            updateMealTable(mealLogData);
            
            // Reset form
            mealLogForm.reset();
        } catch (error) {
            console.error('Error logging meal:', error);
            displayErrorMessage('Failed to log meal. Please try again.');
        }
    });

    function updateMealTable(mealData) {
        const mealList = document.getElementById('mealList');
        const newRow = mealList.insertRow();
        
        newRow.innerHTML = `
            <td>${mealData.mealType}</td>
            <td>${mealData.foodItem}</td>
            <td>${calculateCalories(mealData)}</td>
            <td>${calculateCarbs(mealData)}</td>
        `;
    }
    
    // Placeholder functions for calorie and carb calculations
    function calculateCalories(mealData) {
        // Implement your calorie calculation logic
        return Math.round(mealData.weight * 0.5); // Example calculation
    }
    
    function calculateCarbs(mealData) {
        // Implement your carb calculation logic
        return Math.round(mealData.weight * 0.2); // Example calculation
    }
    
    function displayMealLogSuccess(result) {
        const alertMessage = document.getElementById('alertMessage');
        alertMessage.textContent = result.message;
        alertMessage.classList.add('success');
        
        // Remove success message after 3 seconds
        setTimeout(() => {
            alertMessage.textContent = '';
            alertMessage.classList.remove('success');
        }, 3000);
    }
    
    function displayErrorMessage(message) {
        const alertMessage = document.getElementById('alertMessage');
        alertMessage.textContent = message;
        alertMessage.classList.add('error');
        
        // Remove error message after 3 seconds
        setTimeout(() => {
            alertMessage.textContent = '';
            alertMessage.classList.remove('error');
        }, 3000);
    }


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

    // function logMeal(e) {
    //     e.preventDefault();

    //     const mealType = document.getElementById("mealType").value;
    //     const food_item = document.getElementById("foodItem").value;
    //     const weight = document.getElementById("weight").value;
    //     const proportion = document.getElementById("proportion").value;

    //     const mealData = {
    //         mealType,
    //         food_item,
    //         weight,
    //         proportion,
    //     };

    //     newRow.innerHTML = `
    //         <td>${mealType}</td>
    //         <td>${food_item}</td>
    //         <td>${calories}</td>
    //         <td>${carbs}</td>
    //     `;

    //     fetch("/nutrition/meal/log", {
    //         method: "POST",
    //         headers: {
    //             "Content-Type": "application/json",
    //         },
    //         body: JSON.stringify(mealData),
    //     })
    //         .then((response) => response.json())
    //         .then((data) => {
    //             console.log("Success:", data);
    //             alert("Meal logged successfully!");
    //             document.getElementById("mealLogForm").reset();
    //         })
    //         .catch((error) => {
    //             console.error("Error:", error);
    //             alert("Error logging meal. Please try again.");
    //         });
    // }

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

        // const mealLogForm = document.getElementById('mealLogForm');
        // if (mealLogForm) {
        //     mealLogForm.addEventListener('submit', logMeal);
        // }

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