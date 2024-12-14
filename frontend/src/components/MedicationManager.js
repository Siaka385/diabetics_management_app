window.onload=()=>{
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
    
            addMedicationForm.addEventListener('submit', (event) => {
                event.preventDefault();
                // In a real application, you would send this data to a backend
                alert('Medication added successfully!');
                addMedicationModal.style.display = 'none';
                addMedicationForm.reset();
            });
    
            // Medication tracking buttons
            const medicationItems = document.querySelectorAll('.medication-item');
            medicationItems.forEach(item => {
                const takenBtn = item.querySelector('.btn-taken');
                const skipBtn = item.querySelector('.btn-skip');
    
                takenBtn.addEventListener('click', () => {
                    alert('Medication marked as taken!');
                    // In a real app, this would update the medication history
                });
    
                skipBtn.addEventListener('click', () => {
                    alert('Medication skipped. Please consult your doctor if you regularly miss doses.');
                    // In a real app, this would log the skipped medication
                });
            });
}