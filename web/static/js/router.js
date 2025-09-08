import { checkAuthStatus } from './auth.js';

class Router {
    constructor() {
        this.routes = new Map();
        this.protectedRoutes = new Set(['/dashboard', '/nutrition', '/bloodsugar', '/blog', '/support', '/medication']);
        this.currentRoute = null;
    }

    async start() {
        // Handle initial page load
        await this.handleRoute(window.location.pathname);

        // Handle browser back/forward
        window.addEventListener('popstate', () => {
            this.handleRoute(window.location.pathname);
        });

        // Handle link clicks
        document.addEventListener('click', (e) => {
            if (e.target.matches('a[href^="/"]')) {
                e.preventDefault();
                this.navigate(e.target.getAttribute('href'));
            }
        });
    }

    async navigate(path) {
        history.pushState(null, null, path);
        await this.handleRoute(path);
    }

    async handleRoute(path) {
        const authStatus = await checkAuthStatus();
        const isProtected = this.protectedRoutes.has(path);

        // Redirect logic
        if (isProtected && !authStatus.authenticated) {
            history.replaceState(null, null, '/login');
            path = '/login';
        } else if ((path === '/login' || path === '/signup') && authStatus.authenticated) {
            history.replaceState(null, null, '/dashboard');
            path = '/dashboard';
        }

        this.currentRoute = path;

        // Route to appropriate page
        switch (path) {
            case '/':
                this.renderHomePage();
                break;
            case '/login':
                await this.renderLoginPage();
                break;
            case '/signup':
                await this.renderSignupPage();
                break;
            case '/dashboard':
                await this.renderDashboard();
                break;
            default:
                // For other protected routes, redirect to server
                if (this.protectedRoutes.has(path)) {
                    window.location.href = path;
                } else {
                    this.render404();
                }
        }
    }

    async renderHomePage() {
        const Home = (await import('./pages/home.js')).default;
        Home();
    }

    async renderDashboard() {
        const Dashboard = (await import('./pages/dashboard.js')).default;
        Dashboard();
    }

    async renderLoginPage() {
        const { renderLogin } = await import('./auth/signin.js');
        renderLogin();
    }

    async renderSignupPage() {
        const { renderSignup } = await import('./auth/signup.js');
        renderSignup();
    }

    render404() {
        const app = document.getElementById('app');
        if (!app) return;

        app.innerHTML = `
            <div class="min-h-screen flex items-center justify-center bg-gray-50">
                <div class="text-center">
                    <h1 class="text-6xl font-bold text-gray-900">404</h1>
                    <p class="text-xl text-gray-600 mt-4">Page not found</p>
                    <a href="/" class="mt-6 inline-block bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700">
                        Go Home
                    </a>
                </div>
            </div>
        `;
    }
}

export default Router;