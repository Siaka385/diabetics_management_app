import { signup } from '../auth.js';

export function renderSignup() {
    const app = document.getElementById('app');
    if (!app) return;

    app.innerHTML = `
        <div class="min-h-screen flex items-center justify-center bg-gray-50">
            <div class="max-w-md w-full p-8 bg-white rounded-lg shadow-xl">
                <div class="text-center mb-8">
                    <div class="mx-auto h-12 w-12 flex items-center justify-center rounded-full bg-green-100 mb-4">
                        <span class="text-2xl">🌟</span>
                    </div>
                    <h2 class="text-3xl font-extrabold text-gray-900">
                        Create your account
                    </h2>
                    <p class="mt-2 text-sm text-gray-600">
                        Or
                        <a href="/login" class="font-medium text-blue-600 hover:text-blue-500">
                            sign in to existing account
                        </a>
                    </p>
                </div>
                <form id="signupForm">
                    <div class="mb-4">
                        <label for="username" class="block mb-2 font-medium text-gray-700">Username</label>
                        <input id="username" name="username" type="text" required 
                            class="w-full px-3 py-3 border border-gray-300 rounded-md text-base focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500" 
                            placeholder="Choose a username">
                    </div>
                    <div class="mb-4">
                        <label for="email" class="block mb-2 font-medium text-gray-700">Email address</label>
                        <input id="email" name="email" type="email" required 
                            class="w-full px-3 py-3 border border-gray-300 rounded-md text-base focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500" 
                            placeholder="Enter your email">
                    </div>
                    <div class="mb-4">
                        <label for="password" class="block mb-2 font-medium text-gray-700">Password</label>
                        <input id="password" name="password" type="password" required 
                            class="w-full px-3 py-3 border border-gray-300 rounded-md text-base focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500" 
                            placeholder="Create a password">
                    </div>
                    <div class="mb-4">
                        <label for="confirmPassword" class="block mb-2 font-medium text-gray-700">Confirm Password</label>
                        <input id="confirmPassword" name="confirmPassword" type="password" required 
                            class="w-full px-3 py-3 border border-gray-300 rounded-md text-base focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500" 
                            placeholder="Confirm your password">
                    </div>

                    <div id="errorMessage" class="hidden text-red-600 text-sm text-center mb-4"></div>
                    <div id="successMessage" class="hidden text-green-600 text-sm text-center mb-4"></div>

                    <div class="space-y-3">
                        <button type="submit" 
                            class="w-full px-3 py-3 bg-green-600 text-white border-none rounded-md font-medium cursor-pointer transition-colors duration-200 hover:bg-green-700 disabled:opacity-50 disabled:cursor-not-allowed">
                            Create Account
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
    const form = document.getElementById('signupForm');
    const errorDiv = document.getElementById('errorMessage');
    const successDiv = document.getElementById('successMessage');

    form.addEventListener('submit', async (e) => {
        e.preventDefault();

        const username = document.getElementById('username').value;
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;
        const confirmPassword = document.getElementById('confirmPassword').value;
        const submitBtn = form.querySelector('button[type="submit"]');

        // Hide messages
        errorDiv.classList.add('hidden');
        successDiv.classList.add('hidden');

        // Validate passwords match
        if (password !== confirmPassword) {
            errorDiv.textContent = 'Passwords do not match';
            errorDiv.classList.remove('hidden');
            return;
        }

        // Disable button and show loading
        submitBtn.disabled = true;
        submitBtn.innerHTML = 'Creating Account...';

        try {
            const result = await signup(username, email, password);

            if (result.status === 'success') {
                successDiv.textContent = 'Account created successfully! Redirecting to login...';
                successDiv.classList.remove('hidden');

                // Redirect to login after 2 seconds
                setTimeout(() => {
                    window.history.pushState(null, null, '/login');
                    window.dispatchEvent(new PopStateEvent('popstate'));
                }, 2000);
            } else {
                errorDiv.textContent = result.message || 'Signup failed';
                errorDiv.classList.remove('hidden');
            }
        } catch (error) {
            errorDiv.textContent = 'Network error occurred';
            errorDiv.classList.remove('hidden');
        } finally {
            submitBtn.disabled = false;
            submitBtn.innerHTML = 'Create Account';
        }
    });
}