document.addEventListener('DOMContentLoaded', () => {
    // Modal functionality for adding medication
    const addMedicationBtn = document.getElementById('addMedicationBtn');
    const addMedicationModal = document.getElementById('addMedicationModal');
    const closeModal = document.getElementById('closeModal');
    const addMedicationForm = document.getElementById('addMedicationForm');

    if (addMedicationBtn) {
        addMedicationBtn.addEventListener('click', () => {
            if (addMedicationModal) {
                addMedicationModal.style.display = 'block';
            }
        });
    }

    if (closeModal) {
        closeModal.addEventListener('click', () => {
            if (addMedicationModal) {
                addMedicationModal.style.display = 'none';
            }
        });
    }

    window.addEventListener('click', (event) => {
        if (event.target === addMedicationModal) {
            if (addMedicationModal) {
                addMedicationModal.style.display = 'none';
            }
        }
    });

    if (addMedicationForm) {
        addMedicationForm.addEventListener('submit', async (event) => {
            event.preventDefault();
            const formData = new FormData(addMedicationForm);

            try {
                const response = await fetch('/addmedication', {
                    method: 'POST',
                    body: formData
                });

                if (response.ok) {
                    const result = await response.json();
                    alert('Medication added successfully!');
                    if (addMedicationModal) {
                        addMedicationModal.style.display = 'none';
                    }
                    addMedicationForm.reset();
                } else {
                    const errorText = await response.text();
                    alert(`Failed to add medication: ${errorText}`);
                }
            } catch (error) {
                console.error('Error:', error);
                alert('An error occurred while adding the medication. Please try again.');
            }
        });
    }

    // Medication tracking buttons
    const medicationItems = document.querySelectorAll('.medication-item');
    if (medicationItems) {
        medicationItems.forEach(item => {
            const takenBtn = item.querySelector('.btn-taken');
            const skipBtn = item.querySelector('.btn-skip');

            if (takenBtn) {
                takenBtn.addEventListener('click', () => {
                    alert('Medication marked as taken!');
                });
            }

            if (skipBtn) {
                skipBtn.addEventListener('click', () => {
                    alert('Medication skipped. Please consult your doctor if you regularly miss doses.');
                });
            }
        });
    }
});
