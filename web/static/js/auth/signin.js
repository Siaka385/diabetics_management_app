import { login } from '../auth.js';

export function renderLogin() {
    const app = document.getElementById('app');
    if (!app) return;
    
    app.innerHTML = `
        <div class="min-h-screen flex items-center justify-center bg-gray-50">
            <div class="max-w-md w-full p-8 bg-white rounded-lg shadow-xl">
                <div class="text-center mb-8">
                    <div class="mx-auto h-12 w-12 flex items-center justify-center rounded-full bg-blue-100 mb-4">
                        <span class="text-2xl">🏥</span>
                    </div>
                    <h2 class="text-3xl font-extrabold text-gray-900">
                        Sign in to your account
                    </h2>
                    <p class="mt-2 text-sm text-gray-600">
                        Or
                        <a href="/signup" class="font-medium text-blue-600 hover:text-blue-500">
                            create a new account
                        </a>
                    </p>
                </div>
                <form id="loginForm">
                    <div class="mb-4">
                        <label for="username" class="block mb-2 font-medium text-gray-700">Username</label>
                        <input id="username" name="username" type="text" required 
                            class="w-full px-3 py-3 border border-gray-300 rounded-md text-base focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500" 
                            placeholder="Username">
                    </div>
                    <div class="mb-4">
                        <label for="password" class="block mb-2 font-medium text-gray-700">Password</label>
                        <input id="password" name="password" type="password" required 
                            class="w-full px-3 py-3 border border-gray-300 rounded-md text-base focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500" 
                            placeholder="Password">
                    </div>

                    <div id="errorMessage" class="hidden text-red-600 text-sm text-center mb-4"></div>

                    <div class="space-y-3">
                        <button type="submit" 
                            class="w-full px-3 py-3 bg-blue-600 text-white border-none rounded-md font-medium cursor-pointer transition-colors duration-200 hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed">
                            Sign in
                        </button>
                        <button type="button" onclick="window.history.pushState(null, null, '/'); window.dispatchEvent(new PopStateEvent('popstate'));" 
                            class="w-full px-3 py-3 bg-gray-200 text-gray-700 border-none rounded-md font-medium cursor-pointer transition-colors duration-200 hover:bg-gray-300">
                            Cancel
                        </button>
                    </div>
                </form>
            </div>
        </div>
    `;

    // Add form handler
    const form = document.getElementById('loginForm');
    const errorDiv = document.getElementById('errorMessage');
    
    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        
        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;
        const submitBtn = form.querySelector('button[type="submit"]');
        
        // Disable button and show loading
        submitBtn.disabled = true;
        submitBtn.textContent = 'Signing in...';
        errorDiv.classList.add('hidden');
        
        try {
            const result = await login(username, password);
            
            if (result.status === 'success') {
                // Use client-side navigation
                window.history.pushState(null, null, '/dashboard');
                window.dispatchEvent(new PopStateEvent('popstate'));
            } else {
                errorDiv.textContent = result.message || 'Login failed';
                errorDiv.classList.remove('hidden');
            }
        } catch (error) {
            errorDiv.textContent = 'Network error occurred';
            errorDiv.classList.remove('hidden');
        } finally {
            submitBtn.disabled = false;
            submitBtn.innerHTML = 'Sign in';
        }
    });
}