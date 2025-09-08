export async function renderSupport() {
    try {
        const response = await fetch('/support');
        if (!response.ok) throw new Error('Failed to fetch');
        
        const userData = await response.json();
        const app = document.getElementById('app');
        if (!app) return;
        
        app.innerHTML = `
            <div class="min-h-screen bg-gray-50 flex">
                ${renderSidebar(userData, '/support')}
                <div class="flex-1 p-8">
                    <h1 class="text-3xl font-bold mb-8">Support Groups</h1>
                    
                    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
                        <div class="lg:col-span-2 space-y-6">
                            ${renderSupportGroups()}
                        </div>
                        
                        <div class="space-y-6">
                            <div class="bg-white p-6 rounded-xl shadow-sm border">
                                <h3 class="text-lg font-semibold mb-4">Quick Actions</h3>
                                <div class="space-y-3">
                                    <button class="w-full bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 transition-colors">
                                        Join a Group
                                    </button>
                                    <button class="w-full bg-green-600 text-white py-2 px-4 rounded-lg hover:bg-green-700 transition-colors">
                                        Create Group
                                    </button>
                                    <button class="w-full bg-purple-600 text-white py-2 px-4 rounded-lg hover:bg-purple-700 transition-colors">
                                        Find Local Events
                                    </button>
                                </div>
                            </div>
                            
                            <div class="bg-white p-6 rounded-xl shadow-sm border">
                                <h3 class="text-lg font-semibold mb-4">Community Stats</h3>
                                <div class="space-y-3">
                                    <div class="flex justify-between">
                                        <span class="text-gray-600">Total Members</span>
                                        <span class="font-semibold">2,847</span>
                                    </div>
                                    <div class="flex justify-between">
                                        <span class="text-gray-600">Active Groups</span>
                                        <span class="font-semibold">156</span>
                                    </div>
                                    <div class="flex justify-between">
                                        <span class="text-gray-600">Messages Today</span>
                                        <span class="font-semibold">1,234</span>
                                    </div>
                                </div>
                            </div>
                            
                            <div class="bg-white p-6 rounded-xl shadow-sm border">
                                <h3 class="text-lg font-semibold mb-4">Upcoming Events</h3>
                                <div class="space-y-3">
                                    <div class="p-3 bg-blue-50 rounded-lg">
                                        <div class="font-medium text-sm">Diabetes Support Meetup</div>
                                        <div class="text-xs text-gray-500">Tomorrow 7:00 PM</div>
                                    </div>
                                    <div class="p-3 bg-green-50 rounded-lg">
                                        <div class="font-medium text-sm">Healthy Cooking Workshop</div>
                                        <div class="text-xs text-gray-500">Friday 2:00 PM</div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        `;
        
    } catch (error) {
        window.location.href = '/login';
    }
}

function renderSupportGroups() {
    const groups = [
        {
            name: "Type 1 Diabetes Support",
            members: 245,
            description: "A supportive community for people living with Type 1 diabetes. Share experiences, tips, and encouragement.",
            lastActivity: "2 hours ago",
            category: "Medical Support"
        },
        {
            name: "Newly Diagnosed Support",
            members: 189,
            description: "For those recently diagnosed with diabetes. Get guidance and support from others who understand your journey.",
            lastActivity: "4 hours ago",
            category: "Beginner Support"
        },
        {
            name: "Diabetes Parents Network",
            members: 156,
            description: "Parents supporting parents of children with diabetes. Share resources, experiences, and emotional support.",
            lastActivity: "1 day ago",
            category: "Family Support"
        },
        {
            name: "Fitness & Diabetes",
            members: 298,
            description: "Combining fitness and diabetes management. Share workout routines, tips, and motivation.",
            lastActivity: "3 hours ago",
            category: "Lifestyle"
        }
    ];
    
    return groups.map(group => `
        <div class="bg-white rounded-xl shadow-sm border p-6 hover:shadow-md transition-shadow">
            <div class="flex justify-between items-start mb-4">
                <div>
                    <h3 class="text-xl font-semibold mb-1">${group.name}</h3>
                    <span class="text-sm text-blue-600 bg-blue-100 px-2 py-1 rounded-full">${group.category}</span>
                </div>
                <div class="text-right">
                    <div class="text-sm text-gray-500">${group.members} members</div>
                    <div class="text-xs text-gray-400">Active ${group.lastActivity}</div>
                </div>
            </div>
            <p class="text-gray-600 mb-4">${group.description}</p>
            <div class="flex gap-3">
                <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors text-sm">
                    Join Group
                </button>
                <button class="border border-gray-300 text-gray-700 px-4 py-2 rounded-lg hover:bg-gray-50 transition-colors text-sm">
                    View Details
                </button>
            </div>
        </div>
    `).join('');
}

function renderSidebar(userData, currentPage) {
    return `
        <div class="w-64 bg-white shadow-lg">
            <div class="p-6 border-b">
                <div class="text-center">
                    <div class="w-16 h-16 bg-blue-600 rounded-full flex items-center justify-center text-white text-2xl font-bold mx-auto mb-3">
                        ${userData.Abbrev}
                    </div>
                    <h2 class="text-lg font-semibold">${userData.Name}</h2>
                </div>
            </div>
            <nav class="p-4">
                <ul class="space-y-2">
                    <li><a href="/dashboard" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg ${currentPage === '/dashboard' ? 'text-blue-600 bg-blue-50' : ''}"><span class="mr-3">🏠</span>Dashboard</a></li>
                    <li><a href="/bloodsugar" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg ${currentPage === '/bloodsugar' ? 'text-blue-600 bg-blue-50' : ''}"><span class="mr-3">📊</span>Blood Sugar</a></li>
                    <li><a href="/nutrition" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg ${currentPage === '/nutrition' ? 'text-blue-600 bg-blue-50' : ''}"><span class="mr-3">🥗</span>Nutrition</a></li>
                    <li><a href="/medication" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg ${currentPage === '/medication' ? 'text-blue-600 bg-blue-50' : ''}"><span class="mr-3">💊</span>Medications</a></li>
                    <li><a href="/blog" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg ${currentPage === '/blog' ? 'text-blue-600 bg-blue-50' : ''}"><span class="mr-3">📚</span>Blog</a></li>
                    <li><a href="/support" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg ${currentPage === '/support' ? 'text-blue-600 bg-blue-50' : ''}"><span class="mr-3">🤝</span>Support</a></li>
                </ul>
            </nav>
            <div class="p-4 border-t mt-auto">
                <button onclick="fetch('/auth/signout', {method: 'POST'}).then(() => window.location.href = '/')" 
                    class="w-full flex items-center px-4 py-3 text-red-600 hover:bg-red-50 rounded-lg">
                    <span class="mr-3">🚪</span>Logout
                </button>
            </div>
        </div>
    `;
}