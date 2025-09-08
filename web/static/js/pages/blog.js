export async function renderBlog() {
    try {
        const response = await fetch('/blog');
        if (!response.ok) throw new Error('Failed to fetch');
        
        const userData = await response.json();
        const app = document.getElementById('app');
        if (!app) return;
        
        app.innerHTML = `
            <div class="min-h-screen bg-gray-50 flex">
                ${renderSidebar(userData, '/blog')}
                <div class="flex-1 p-8">
                    <div class="flex justify-between items-center mb-8">
                        <h1 class="text-3xl font-bold">DiaWise Blog</h1>
                        <button id="createTopicBtn" class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
                            Create New Topic
                        </button>
                    </div>
                    
                    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
                        <div class="lg:col-span-2 space-y-6">
                            ${renderBlogPosts(userData.Posts)}
                        </div>
                        
                        <div class="bg-white p-6 rounded-xl shadow-sm border">
                            <h3 class="text-lg font-semibold mb-4">Active Discussions</h3>
                            <div class="space-y-3">
                                <div class="flex items-center p-3 bg-gray-50 rounded-lg">
                                    <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center mr-3">
                                        <span>💬</span>
                                    </div>
                                    <div>
                                        <div class="font-medium text-sm">Medication Side Effects</div>
                                        <div class="text-xs text-gray-500">24 new messages</div>
                                    </div>
                                </div>
                                <div class="flex items-center p-3 bg-gray-50 rounded-lg">
                                    <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center mr-3">
                                        <span>🩺</span>
                                    </div>
                                    <div>
                                        <div class="font-medium text-sm">Insulin Pump Questions</div>
                                        <div class="text-xs text-gray-500">18 new messages</div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            
            <!-- Create Topic Modal -->
            <div id="createTopicModal" class="fixed inset-0 bg-black bg-opacity-50 hidden items-center justify-center z-50">
                <div class="bg-white p-6 rounded-xl max-w-md w-full mx-4">
                    <div class="flex justify-between items-center mb-4">
                        <h2 class="text-xl font-semibold">Create New Topic</h2>
                        <button id="closeModal" class="text-gray-500 hover:text-gray-700">&times;</button>
                    </div>
                    <form id="createTopicForm" class="space-y-4">
                        <div>
                            <label class="block mb-2 font-medium text-gray-700">Topic Title</label>
                            <input type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:border-blue-500" placeholder="Enter topic title">
                        </div>
                        <div>
                            <label class="block mb-2 font-medium text-gray-700">Description</label>
                            <textarea rows="4" required class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:border-blue-500" placeholder="Describe your topic..."></textarea>
                        </div>
                        <button type="submit" class="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 transition-colors">
                            Create Topic
                        </button>
                    </form>
                </div>
            </div>
        `;
        
        // Add event handlers
        document.getElementById('createTopicBtn').addEventListener('click', () => {
            document.getElementById('createTopicModal').classList.remove('hidden');
            document.getElementById('createTopicModal').classList.add('flex');
        });
        
        document.getElementById('closeModal').addEventListener('click', () => {
            document.getElementById('createTopicModal').classList.add('hidden');
            document.getElementById('createTopicModal').classList.remove('flex');
        });
        
        document.getElementById('createTopicForm').addEventListener('submit', (e) => {
            e.preventDefault();
            alert('Topic created successfully!');
            document.getElementById('createTopicModal').classList.add('hidden');
            document.getElementById('createTopicModal').classList.remove('flex');
        });
        
        // Add click handlers for blog posts
        document.querySelectorAll('[data-post-id]').forEach(postElement => {
            postElement.addEventListener('click', (e) => {
                e.preventDefault();
                const postId = e.currentTarget.getAttribute('data-post-id');
                if (window.router && window.router.navigate) {
                    window.router.navigate(`/post/${postId}`);
                } else {
                    // Fallback: trigger router manually
                    history.pushState(null, null, `/post/${postId}`);
                    window.dispatchEvent(new PopStateEvent('popstate'));
                }
            });
        });
        
    } catch (error) {
        window.location.href = '/login';
    }
}

function renderBlogPosts(posts) {
    if (!posts || posts.length === 0) {
        return '<div class="text-center text-gray-500">No posts available</div>';
    }
    
    return posts.map(post => `
        <div class="bg-white rounded-xl shadow-sm border overflow-hidden hover:shadow-md transition-shadow cursor-pointer" data-post-id="${post.ID}">
            <img src="/static/images/${post.Image}" alt="${post.Title}" class="w-full h-48 object-cover">
            <div class="p-6">
                <h3 class="text-xl font-semibold mb-2 hover:text-blue-600">${post.Title}</h3>
                <div class="text-sm text-gray-500 mb-3">
                    By ${post.Author} • ${post.Date}
                </div>
                <div class="text-gray-600 mb-4">${post.Excerpt}</div>
                <button class="text-blue-600 hover:text-blue-700 font-medium read-more-btn">Read More →</button>
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
                    <li><a href="/education" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg ${currentPage === '/education' ? 'text-blue-600 bg-blue-50' : ''}"><span class="mr-3">📖</span>Education</a></li>
                    <li><a href="/glucose-tracker" class="flex items-center px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-lg ${currentPage === '/glucose-tracker' ? 'text-blue-600 bg-blue-50' : ''}"><span class="mr-3">📈</span>Glucose Tracker</a></li>
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