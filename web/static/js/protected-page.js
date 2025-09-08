// Protected page initialization
class ProtectedPage {
    static init(pageName = '') {
        // Require authentication
        if (!Auth.requireAuth()) {
            return;
        }
        
        // Initialize navigation
        Navigation.init(window.location.pathname);
        
        // Add page-specific functionality
        this.setupPageFeatures(pageName);
    }
    
    static setupPageFeatures(pageName) {
        // Add common protected page features
        this.addUserInfo();
        this.setupLogoutHandlers();
    }
    
    static addUserInfo() {
        // Add user info display if needed
        const userInfoElements = document.querySelectorAll('.user-info');
        userInfoElements.forEach(element => {
            // You can populate user info here if available
            element.style.display = 'block';
        });
    }
    
    static setupLogoutHandlers() {
        // Setup logout buttons
        const logoutButtons = document.querySelectorAll('.logout-btn, [data-logout]');
        logoutButtons.forEach(button => {
            button.addEventListener('click', (e) => {
                e.preventDefault();
                Auth.logout();
            });
        });
    }
}

// Auto-initialize for protected pages
document.addEventListener('DOMContentLoaded', function() {
    // Check if this is a protected page
    const protectedRoutes = ['/dashboard', '/nutrition', '/bloodsugar', '/blog', '/support', '/medication', '/addmedication'];
    const currentPath = window.location.pathname;
    
    if (protectedRoutes.includes(currentPath)) {
        ProtectedPage.init(currentPath);
    }
});