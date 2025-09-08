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
        try {
            await fetch('/logout', { method: 'GET' });
            window.location.href = '/login';
        } catch (error) {
            window.location.href = '/login';
        }
    }
    
    static async isAuthenticated() {
        try {
            const response = await fetch('/auth/status');
            const data = await response.json();
            return data.authenticated;
        } catch (error) {
            return false;
        }
    }
    
    static async requireAuth() {
        const isAuth = await this.isAuthenticated();
        if (!isAuth) {
            window.location.href = '/login';
            return false;
        }
        return true;
    }
    
    static async redirectIfAuthenticated() {
        const isAuth = await this.isAuthenticated();
        if (isAuth) {
            window.location.href = '/dashboard';
            return true;
        }
        return false;
    }
}