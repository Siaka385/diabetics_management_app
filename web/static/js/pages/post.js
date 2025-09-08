export async function renderPost(postId) {
    try {
        const response = await fetch(`/post/${postId}`);
        if (!response.ok) throw new Error('Failed to fetch');
        
        const data = await response.json();
        const app = document.getElementById('app');
        if (!app) return;
        
        app.innerHTML = `
            <div class="min-h-screen bg-gray-50 flex">
                ${renderSidebar(data, data.CurrentPage)}
                <div class="flex-1 p-8">
                    <div class="max-w-4xl mx-auto">
                        <button onclick="history.back()" class="mb-6 flex items-center text-blue-600 hover:text-blue-700">
                            <span class="mr-2">←</span>Back to Blog
                        </button>
                        
                        <article class="bg-white rounded-xl shadow-sm border overflow-hidden">
                            <img src="/static/images/${data.Post.Image}" alt="${data.Post.Title}" class="w-full h-64 object-cover">
                            
                            <div class="p-8">
                                <h1 class="text-3xl font-bold mb-4">${data.Post.Title}</h1>
                                
                                <div class="flex items-center text-gray-600 mb-6">
                                    <span class="mr-4">By ${data.Post.Author}</span>
                                    <span>${data.Post.Date}</span>
                                </div>
                                
                                <div class="prose max-w-none">
                                    ${data.Post.Content}
                                </div>
                            </div>
                        </article>
                    </div>
                </div>
            </div>
        `;
        
    } catch (error) {
        window.location.href = '/login';
    }
}

function renderSidebar(userData, currentPage) {
    return `
        <div class="w-64 bg-white shadow-lg">
            <div class="p-6 border-b">
                <div class="text-center">
                    <div class="w-16 h-16 bg-blue-600 rounded-full flex items-center justify-center text-white text-2xl font-bold mx-auto mb-3">
                        ${userData.Abbrev || userData.name?.charAt(0) || 'U'}
                    </div>
                    <h2 class="text-lg font-semibold">${userData.Name || userData.name || 'User'}</h2>
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