// Modal functionality for adding medication
const addMedicationBtn = document.getElementById('addMedicationBtn');
const addMedicationModal = document.getElementById('addMedicationModal');
const closeModal = document.getElementById('closeModal');
const addMedicationForm = document.getElementById('addMedicationForm');

addMedicationBtn.addEventListener('click', () => {
    addMedicationModal.style.display = 'block';
});

closeModal.addEventListener('click', () => {
    addMedicationModal.style.display = 'none';
});

window.addEventListener('click', (event) => {
    if (event.target === addMedicationModal) {
        addMedicationModal.style.display = 'none';
    }
});

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
            addMedicationModal.style.display = 'none';
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

// Medication tracking buttons
const medicationItems = document.querySelectorAll('.medication-item');
medicationItems.forEach(item => {
    const takenBtn = item.querySelector('.btn-taken');
    const skipBtn = item.querySelector('.btn-skip');

    takenBtn.addEventListener('click', () => {
        alert('Medication marked as taken!');
    });

    skipBtn.addEventListener('click', () => {
        alert('Medication skipped. Please consult your doctor if you regularly miss doses.');
    });
});