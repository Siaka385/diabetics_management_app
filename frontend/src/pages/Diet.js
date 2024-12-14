
        //DIETs AND NUTRITION
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

        // Meal Plan Generation Logic
        function generateMealPlan(e) {
            e.preventDefault();
            
            const selectedDuration = document.querySelector('#mealPlanDurationTags .selected')?.dataset.value;
            const selectedMealType = document.querySelector('#mealTypeTags .selected')?.dataset.value;
            const selectedDietPref = document.querySelector('#dietPreferenceTags .selected')?.dataset.value;
            const preferredFoods = document.getElementById('preferredFoods').value.split(',').map(f => f.trim());
            const avoidFoods = document.getElementById('avoidFoods').value.split(',').map(f => f.trim());

            const mealPlanContainer = document.getElementById('generatedMealPlan');
            
            // Simulated Meal Plan Generation with Local Alternatives
            const mealPlans = {
                'low-carb': {
                    breakfast: {
                        main: 'Scrambled Eggs with Spinach',
                        local_alt: 'Tofu Scramble with Local Greens',
                        carbs: 3
                    },
                    lunch: {
                        main: 'Grilled Chicken Salad',
                        local_alt: 'Grilled Fish with Mixed Vegetable Salad',
                        carbs: 10
                    },
                    dinner: {
                        main: 'Baked Salmon with Roasted Vegetables',
                        local_alt: 'Steamed Local Fish with Stir-fried Vegetables',
                        carbs: 8
                    }
                },
                // Add more diet preference meal plans
            };

            let generatedPlan = `<h3>Personalized Meal Plan</h3>`;
            
            if (selectedDietPref && mealPlans[selectedDietPref]) {
                const plan = mealPlans[selectedDietPref];
                
                for (const [meal, details] of Object.entries(plan)) {
                    generatedPlan += `
                        <div class="meal-plan-entry">
                            <h4>${meal.charAt(0).toUpperCase() + meal.slice(1)}</h4>
                            <p><strong>Recommended:</strong> ${details.main}</p>
                            <p><strong>Local Alternative:</strong> ${details.local_alt}</p>
                            <p><strong>Carbohydrates:</strong> ${details.carbs}g</p>
                        </div>
                    `;
                }
            } else {
                generatedPlan += `<p>Unable to generate a meal plan. Please select preferences.</p>`;
            }

            mealPlanContainer.innerHTML = generatedPlan;
        }

        // Initialize page
            setupPreferenceTags();
            
            const mealPlanForm = document.getElementById('mealPlanPreferencesForm');
            mealPlanForm.addEventListener('submit', generateMealPlan);
        
