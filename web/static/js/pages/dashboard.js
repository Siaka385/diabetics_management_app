const Dashboard = async () => {
    try {
        const response = await fetch('/dashboard');
        if (response.ok) {
            const userData = await response.json();
            const app = document.getElementById('app');
            if (!app) return;

            app.innerHTML = `
                    <div class="min-h-screen bg-gray-50 flex">
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
                                    <li><a href="/dashboard" class="flex items-center px-4 py-3 text-blue-600 bg-blue-50 rounded-lg"><span class="mr-3">🏠</span>Dashboard</a></li>
                                    <li><a href="/bloodsugar" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg"><span class="mr-3">📊</span>Blood Sugar</a></li>
                                    <li><a href="/nutrition" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg"><span class="mr-3">🥗</span>Nutrition</a></li>
                                    <li><a href="/medication" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg"><span class="mr-3">💊</span>Medications</a></li>
                                    <li><a href="/blog" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg"><span class="mr-3">📚</span>Blog</a></li>
                                    <li><a href="/support" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg"><span class="mr-3">🤝</span>Support</a></li>
                                </ul>
                            </nav>
                            <div class="p-4 border-t mt-auto">
                                <button onclick="fetch('/auth/signout', {method: 'POST'}).then(() => window.location.href = '/')" class="w-full flex items-center px-4 py-3 text-red-600 hover:bg-red-50 rounded-lg">
                                    <span class="mr-3">🚪</span>Logout
                                </button>
                            </div>
                        </div>
                        <div class="flex-1 p-8">
                            <h1 class="text-3xl font-bold mb-8">Welcome back, ${userData.Name}!</h1>
                            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
                                <div class="bg-white p-6 rounded-xl shadow-sm border">
                                    <h3 class="text-lg font-semibold mb-4">Blood Sugar</h3>
                                    <div class="text-2xl font-bold text-gray-900">128 <span class="text-sm text-gray-500">mg/dL</span></div>
                                    <div class="w-full bg-gray-200 rounded-full h-2 mt-3">
                                        <div class="bg-green-500 h-2 rounded-full" style="width: 65%"></div>
                                    </div>
                                </div>
                                <div class="bg-white p-6 rounded-xl shadow-sm border">
                                    <h3 class="text-lg font-semibold mb-4">Nutrition</h3>
                                    <div class="text-2xl font-bold text-gray-900">1800 <span class="text-sm text-gray-500">cal</span></div>
                                    <div class="w-full bg-gray-200 rounded-full h-2 mt-3">
                                        <div class="bg-blue-500 h-2 rounded-full" style="width: 75%"></div>
                                    </div>
                                </div>
                                <div class="bg-white p-6 rounded-xl shadow-sm border">
                                    <h3 class="text-lg font-semibold mb-4">Medications</h3>
                                    <div class="text-2xl font-bold text-gray-900">2 <span class="text-sm text-gray-500">pending</span></div>
                                    <div class="w-full bg-gray-200 rounded-full h-2 mt-3">
                                        <div class="bg-yellow-500 h-2 rounded-full" style="width: 85%"></div>
                                    </div>
                                </div>
                                <div class="bg-white p-6 rounded-xl shadow-sm border">
                                    <h3 class="text-lg font-semibold mb-4">Support</h3>
                                    <div class="text-2xl font-bold text-gray-900">12 <span class="text-sm text-gray-500">members</span></div>
                                    <div class="w-full bg-gray-200 rounded-full h-2 mt-3">
                                        <div class="bg-purple-500 h-2 rounded-full" style="width: 55%"></div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                `;
        } else {
            this.navigate('/login');
        }
    } catch (error) {
        this.navigate('/login');
    }
};

export default Dashboard;