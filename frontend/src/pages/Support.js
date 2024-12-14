
    /*Support and community*/
       // Modal functionality
       const createTopicBtn = document.getElementById('createTopicBtn');
        const createTopicModal = document.getElementById('createTopicModal');
        const closeModal = document.getElementById('closeModal');
        const createTopicForm = document.getElementById('createTopicForm');

        createTopicBtn.addEventListener('click', () => {
            createTopicModal.style.display = 'block';
        });

        closeModal.addEventListener('click', () => {
            createTopicModal.style.display = 'none';
        });

        window.addEventListener('click', (event) => {
            if (event.target === createTopicModal) {
                createTopicModal.style.display = 'none';
            }
        });

        createTopicForm.addEventListener('submit', (event) => {
            event.preventDefault();
            // In a real application, you would send this data to a backend
            alert('Topic created successfully!');
            createTopicModal.style.display = 'none';
            createTopicForm.reset();
        });
