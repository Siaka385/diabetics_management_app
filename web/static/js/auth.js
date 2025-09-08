// Frontend Authentication Utility
class Auth {
    static async login(username, password) {
        try {
            const response = await fetch('/auth/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ username, password })
            });
            
            const data = await response.json();
            
            if (data.status === 'success') {
                window.location.href = data.redirect || '/dashboard';
                return { success: true };
            } else {
                return { success: false, message: data.message };
            }
        } catch (error) {
            return { success: false, message: 'Network error occurred' };
        }
    }
    
    static async logout() {
        window.location.href = '/logout';
    }
    
    static checkAuth() {
        // Check if user is on a protected route without auth
        const protectedRoutes = ['/dashboard', '/nutrition', '/bloodsugar', '/blog', '/support', '/medication'];
        const currentPath = window.location.pathname;
        
        if (protectedRoutes.includes(currentPath)) {
            // The server will handle redirect if not authenticated
            return true;
        }
        return false;
    }
}

// Auto-check auth on page load
document.addEventListener('DOMContentLoaded', function() {
    Auth.checkAuth();
});