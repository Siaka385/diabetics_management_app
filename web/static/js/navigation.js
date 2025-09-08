// Navigation utility for authenticated pages
class Navigation {
    static createNavBar(currentPage = '') {
        const navItems = [
            { path: '/dashboard', label: 'Dashboard', icon: '🏠' },
            { path: '/nutrition', label: 'Nutrition', icon: '🥗' },
            { path: '/bloodsugar', label: 'Blood Sugar', icon: '📊' },
            { path: '/medication', label: 'Medication', icon: '💊' },
            { path: '/blog', label: 'Community', icon: '👥' },
            { path: '/support', label: 'Support', icon: '🆘' }
        ];

        const nav = document.createElement('nav');
        nav.className = 'main-nav';
        nav.innerHTML = `
            <div class="nav-brand">
                <a href="/dashboard">Diawise</a>
            </div>
            <ul class="nav-links">
                ${navItems.map(item => `
                    <li class="${currentPage === item.path ? 'active' : ''}">
                        <a href="#" onclick="Navigation.navigate('${item.path}')">
                            <span class="nav-icon">${item.icon}</span>
                            ${item.label}
                        </a>
                    </li>
                `).join('')}
            </ul>
            <div class="nav-user">
                <button onclick="Auth.logout()" class="logout-btn">Logout</button>
            </div>
        `;

        return nav;
    }

    static navigate(path) {
        window.AppRouter.navigate(path);
    }

    static init(currentPage = '') {
        // Navigation is handled by existing sidebar
        // Just add logout functionality to existing buttons
        this.setupLogoutButtons();
    }
    
    static setupLogoutButtons() {
        const logoutButtons = document.querySelectorAll('.logout-btn, [data-logout]');
        logoutButtons.forEach(button => {
            button.addEventListener('click', (e) => {
                e.preventDefault();
                Auth.logout();
            });
        });
    }

    static addNavStyles() {
        if (document.getElementById('nav-styles')) return;
        
        const style = document.createElement('style');
        style.id = 'nav-styles';
        style.textContent = `
            .main-nav {
                display: flex;
                justify-content: space-between;
                align-items: center;
                padding: 1rem 2rem;
                background: #fff;
                border-bottom: 1px solid #e5e7eb;
                box-shadow: 0 1px 3px rgba(0,0,0,0.1);
            }
            .nav-brand a {
                font-size: 1.5rem;
                font-weight: 700;
                color: #2563eb;
                text-decoration: none;
            }
            .nav-links {
                display: flex;
                list-style: none;
                margin: 0;
                padding: 0;
                gap: 2rem;
            }
            .nav-links a {
                display: flex;
                align-items: center;
                gap: 0.5rem;
                padding: 0.5rem 1rem;
                text-decoration: none;
                color: #6b7280;
                border-radius: 0.5rem;
                transition: all 0.2s;
            }
            .nav-links a:hover, .nav-links .active a {
                background: #f3f4f6;
                color: #2563eb;
            }
            .logout-btn {
                padding: 0.5rem 1rem;
                background: #ef4444;
                color: white;
                border: none;
                border-radius: 0.5rem;
                cursor: pointer;
                transition: background 0.2s;
            }
            .logout-btn:hover {
                background: #dc2626;
            }
        `;
        document.head.appendChild(style);
    }
}