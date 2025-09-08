document.addEventListener('DOMContentLoaded', () => {
    function setupMobileMenu() {
        // Create mobile menu toggle
        const mobileMenuToggle = document.createElement('button');
        mobileMenuToggle.classList.add('mobile-menu-toggle');
        mobileMenuToggle.setAttribute('aria-label', 'Toggle Menu');

        // Create menu toggle icon
        const menuToggleIcon = document.createElement('span');
        menuToggleIcon.classList.add('mobile-menu-toggle-icon');
        mobileMenuToggle.appendChild(menuToggleIcon);

        const sidebar = document.querySelector('.sidebar');
        const mainContent = document.querySelector('.main-content');
        
        // Create overlay
        const overlay = document.createElement('div');
        overlay.classList.add('overlay');

        // Insert menu toggle and overlay
        document.body.insertBefore(mobileMenuToggle, document.body.firstChild);
        document.body.insertBefore(overlay, document.body.firstChild);

        // Toggle menu function
        function toggleMenu(event) {
            event.stopPropagation();
            sidebar.classList.toggle('open');
            overlay.classList.toggle('active');
            mobileMenuToggle.classList.toggle('active');
            
            // Prevent body scrolling when menu is open
            document.body.style.overflow = sidebar.classList.contains('open') ? 'hidden' : '';
        }

        // Close menu function
        function closeMenu() {
            sidebar.classList.remove('open');
            overlay.classList.remove('active');
            mobileMenuToggle.classList.remove('active');
            document.body.style.overflow = '';
        }

        // Add event listeners
        mobileMenuToggle.addEventListener('click', toggleMenu);
        overlay.addEventListener('click', closeMenu);

        // Close menu when a sidebar link is clicked
        const sidebarLinks = sidebar.querySelectorAll('.sidebar-nav a');
        sidebarLinks.forEach(link => {
            link.addEventListener('click', closeMenu);
        });

        // Handle window resize
        function handleResize() {
            if (window.innerWidth > 768) {
                closeMenu();
            }
        }

        window.addEventListener('resize', handleResize);
    }

    // Call mobile menu setup
    setupMobileMenu();
});
