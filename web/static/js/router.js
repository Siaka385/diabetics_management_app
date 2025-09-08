// Frontend Router with Authentication
class Router {
    constructor() {
        this.routes = new Map();
        this.protectedRoutes = new Set();
        this.currentRoute = null;
        this.init();
    }

    // Define routes and their protection status
    addRoute(path, handler, isProtected = false) {
        this.routes.set(path, handler);
        if (isProtected) {
            this.protectedRoutes.add(path);
        }
    }

    // Check if user is authenticated by calling auth status endpoint
    async isAuthenticated() {
        try {
            const response = await fetch('/auth/status');
            const data = await response.json();
            return data.authenticated;
        } catch (error) {
            console.log('Auth check failed:', error);
            return false;
        }
    }

    // Navigate to a route
    navigate(path) {
        // Check if route is protected and user is not authenticated
        if (this.protectedRoutes.has(path) && !this.isAuthenticated()) {
            window.location.href = '/login';
            return;
        }

        // If trying to access login/signup while authenticated, redirect to dashboard
        if ((path === '/login' || path === '/signup') && this.isAuthenticated()) {
            window.location.href = '/dashboard';
            return;
        }

        // Navigate to the route
        window.location.href = path;
    }

    // Initialize router
    init() {
        // Define all routes
        this.addRoute('/', () => {}, false);
        this.addRoute('/login', () => {}, false);
        this.addRoute('/signup', () => {}, false);
        this.addRoute('/dashboard', () => {}, true);
        this.addRoute('/nutrition', () => {}, true);
        this.addRoute('/bloodsugar', () => {}, true);
        this.addRoute('/blog', () => {}, true);
        this.addRoute('/support', () => {}, true);
        this.addRoute('/medication', () => {}, true);
        this.addRoute('/addmedication', () => {}, true);
        this.addRoute('/education', () => {}, false);

        // Check current route on page load
        this.checkCurrentRoute();
    }

    // Check if current route is valid and user has access
    async checkCurrentRoute() {
        const currentPath = window.location.pathname;
        this.currentRoute = currentPath;
        const isAuth = await this.isAuthenticated();
        const isProtected = this.protectedRoutes.has(currentPath);
        
        console.log('Router check - path:', currentPath, 'isAuth:', isAuth, 'isProtected:', isProtected);

        // If on protected route without auth, redirect to login
        if (isProtected && !isAuth) {
            console.log('Redirecting to login - protected route without auth');
            window.location.href = '/login';
            return;
        }

        // If on login/signup while authenticated, redirect to dashboard
        if ((currentPath === '/login' || currentPath === '/signup') && isAuth) {
            console.log('Redirecting to dashboard - already authenticated');
            window.location.href = '/dashboard';
            return;
        }
        
        console.log('Router check complete - no redirect needed');
    }

    // Get current authentication status
    getAuthStatus() {
        return {
            isAuthenticated: this.isAuthenticated(),
            currentRoute: this.currentRoute,
            isProtectedRoute: this.protectedRoutes.has(this.currentRoute)
        };
    }
}

// Create global router instance
window.AppRouter = new Router();