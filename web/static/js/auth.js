// Authentication utilities
async function checkAuthStatus() {
    try {
        const response = await fetch('/auth/status');
        const data = await response.json();
        return data;
    } catch (error) {
        return { authenticated: false };
    }
}

async function login(username, password) {
    try {
        const response = await fetch('/auth/signin', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password })
        });
        
        const data = await response.json();
        return data;
    } catch (error) {
        return { status: 'error', message: 'Network error occurred' };
    }
}

async function signup(username, email, password) {
    try {
        const response = await fetch('/auth/signup', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, email, password })
        });
        
        const data = await response.json();
        return data;
    } catch (error) {
        return { status: 'error', message: 'Network error occurred' };
    }
}

async function logout() {
    try {
        await fetch('/auth/signout', { method: 'POST' });
        window.location.href = '/';
    } catch (error) {
        window.location.href = '/';
    }
}

export {
    checkAuthStatus,
    login,
    signup,
    logout
};