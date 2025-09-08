import Router from './router.js';
import { checkAuthStatus } from './auth.js';

class App {
    constructor() {
        this.router = null;
        this.init();
    }

    async init() {
        // Wait for DOM to be ready
        if (document.readyState === 'loading') {
            document.addEventListener('DOMContentLoaded', () => this.start());
        } else {
            this.start();
        }
    }

    async start() {
        // Initialize router
        this.router = new Router();
        
        // Check authentication status
        const authStatus = await checkAuthStatus();
        
        // Set global auth status
        window.authStatus = authStatus;
        
        // Start the router
        this.router.start();
    }
}

// Initialize app
new App();