// Dashboard Client-Side Rendering
class DashboardRenderer {
    constructor() {
        this.userData = null;
        this.init();
    }

    async init() {
        await this.fetchDashboardData();
        this.renderSidebar();
        this.renderMainContent();
    }

    async fetchDashboardData() {
        try {
            const response = await fetch('/dashboard');
            if (response.ok) {
                this.userData = await response.json();
            } else {
                console.error('Failed to fetch dashboard data');
                // Redirect to login if unauthorized
                window.location.href = '/login';
            }
        } catch (error) {
            console.error('Error fetching dashboard data:', error);
        }
    }

    renderSidebar() {
        if (!this.userData) return;

        const sidebarHTML = `
            <aside class="sidebar">
                <div class="profile">
                    <div class="profile-image">${this.userData.Abbrev}</div>
                    <h2>${this.userData.Name}</h2>
                    <div class="edit-profile">✎</div>
                </div>
                <nav>
                    <ul class="sidebar-nav">
                        <li><a href="/dashboard" class="${this.userData.CurrentPage === '/dashboard' ? 'active' : ''}"><span class="nav-icon">🏠</span><span class="nav-text">Dashboard</span></a></li>
                        <li><a href="/bloodsugar" class="${this.userData.CurrentPage === '/bloodsugar' ? 'active' : ''}"><span class="nav-icon">📊</span><span class="nav-text">Blood Sugar</span></a></li>
                        <li><a href="/nutrition" class="${this.userData.CurrentPage === '/nutrition' ? 'active' : ''}"><span class="nav-icon">🥗</span><span class="nav-text">Diet & Nutrition</span></a></li>
                        <li><a href="/medication" class="${this.userData.CurrentPage === '/medication' ? 'active' : ''}"><span class="nav-icon">💊</span><span class="nav-text">Medications</span></a></li>
                        <li><a href="/blog" class="${this.userData.CurrentPage === '/blog' ? 'active' : ''}"><span class="nav-icon">👥</span><span class="nav-text">Education</span></a></li>
                        <li><a href="/support" class="${this.userData.CurrentPage === '/support' ? 'active' : ''}"><span class="nav-icon">🆘</span><span class="nav-text">Support Groups</span></a></li>
                    </ul>
                    <div class="sidebar-logout">
                        <button onclick="Auth.logout()" class="logout-btn"><span class="nav-icon">🚪</span><span class="nav-text">Logout</span></button>
                    </div>
                </nav>
            </aside>
        `;

        // Replace or insert the sidebar
        const existingSidebar = document.querySelector('.sidebar');
        if (existingSidebar) {
            existingSidebar.outerHTML = sidebarHTML;
        } else {
            document.querySelector('.dashboard').insertAdjacentHTML('afterbegin', sidebarHTML);
        }
    }

    renderMainContent() {
        if (!this.userData) return;

        const mainContentHTML = `
            <main class="main-content">
                <div class="header">
                    <h1>Welcome, ${this.userData.Name}</h1>
                    <div class="notifications">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"></path>
                            <path d="M13.73 21a2 2 0 0 1-3.46 0"></path>
                        </svg>
                        <div class="notification-icon">3</div>
                    </div>
                </div>

                <div class="dashboard-grid">
                    <div class="dashboard-card">
                        <div class="card-header">
                            <h3>Blood Sugar</h3>
                            <a href="/bloodsugar">Track</a>
                        </div>
                        <div class="quick-stats">
                            <div class="stat">
                                <span class="stat-value">128</span>
                                <span class="stat-label">mg/dL</span>
                            </div>
                            <div class="stat">
                                <span class="stat-value">5.8</span>
                                <span class="stat-label">HbA1c</span>
                            </div>
                        </div>
                        <div class="progress-bar">
                            <div class="progress" style="width: 65%"></div>
                        </div>
                    </div>

                    <div class="dashboard-card">
                        <div class="card-header">
                            <h3>Diet & Nutrition</h3>
                            <a href="/nutrition">Log Meal</a>
                        </div>
                        <div class="quick-stats">
                            <div class="stat">
                                <span class="stat-value">1800</span>
                                <span class="stat-label">Calories</span>
                            </div>
                            <div class="stat">
                                <span class="stat-value">45g</span>
                                <span class="stat-label">Carbs</span>
                            </div>
                        </div>
                        <div class="progress-bar">
                            <div class="progress" style="width: 75%"></div>
                        </div>
                    </div>

                    <div class="dashboard-card">
                        <div class="card-header">
                            <h3>Medications</h3>
                            <a href="/medication">Manage</a>
                        </div>
                        <div class="quick-stats">
                            <div class="stat">
                                <span class="stat-value">2</span>
                                <span class="stat-label">Pending</span>
                            </div>
                            <div class="stat">
                                <span class="stat-value">3</span>
                                <span class="stat-label">Taken</span>
                            </div>
                        </div>
                        <div class="progress-bar">
                            <div class="progress" style="width: 85%"></div>
                        </div>
                    </div>

                    <div class="dashboard-card">
                        <div class="card-header">
                            <h3>Support Group</h3>
                            <a href="/support">Join</a>
                        </div>
                        <div class="quick-stats">
                            <div class="stat">
                                <span class="stat-value">12</span>
                                <span class="stat-label">Members</span>
                            </div>
                            <div class="stat">
                                <span class="stat-value">3</span>
                                <span class="stat-label">New Posts</span>
                            </div>
                        </div>
                        <div class="progress-bar">
                            <div class="progress" style="width: 55%"></div>
                        </div>
                    </div>
                </div>
            </main>
        `;

        // Replace or insert the main content
        const existingMain = document.querySelector('.main-content');
        if (existingMain) {
            existingMain.outerHTML = mainContentHTML;
        } else {
            document.querySelector('.welcomepage').innerHTML = mainContentHTML;
        }
    }
}

// Initialize dashboard renderer when DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    new DashboardRenderer();
});
