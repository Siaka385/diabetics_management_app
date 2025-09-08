const Home = async () => {
    const app = document.getElementById('app');
    if (!app) return;

    app.innerHTML = `
            <nav class="bg-white shadow-lg sticky top-0 z-50">
                <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                    <div class="flex justify-between h-16">
                        <div class="flex items-center">
                            <h1 class="text-2xl font-bold text-blue-600">DiaWise</h1>
                        </div>
                        <div class="hidden md:flex items-center space-x-8">
                            <a href="#features" class="text-gray-700 hover:text-blue-600">Features</a>
                            <a href="#community" class="text-gray-700 hover:text-blue-600">Community</a>
                            <a href="/login" class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700">Login</a>
                        </div>
                    </div>
                </div>
            </nav>

            <section class="bg-gradient-to-br from-blue-50 to-indigo-100 py-20">
                <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                    <div class="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
                        <div class="text-center lg:text-left">
                            <h1 class="text-4xl lg:text-6xl font-bold text-gray-900 mb-6">
                                Take Control of Your <span class="text-blue-600">Diabetes Journey</span>
                            </h1>
                            <p class="text-xl text-gray-600 mb-8">
                                Join thousands managing diabetes through personalized support, expert guidance, and community.
                            </p>
                            <div class="flex flex-col sm:flex-row gap-4 justify-center lg:justify-start">
                                <a href="/login" class="bg-blue-600 text-white px-8 py-3 rounded-lg text-lg font-semibold hover:bg-blue-700">
                                    Log In
                                </a>
                                <a href="/signup" class="bg-white text-blue-600 border-2 border-blue-600 px-8 py-3 rounded-lg text-lg font-semibold hover:bg-blue-50">
                                    Sign Up
                                </a>
                            </div>
                        </div>
                        <div class="text-center">
                            <img src="/static/images/hero-image.jpg" alt="Diabetes Management" class="rounded-lg shadow-2xl max-w-full h-auto">
                        </div>
                    </div>
                </div>
            </section>

            <section id="features" class="py-20 bg-white">
                <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                    <div class="text-center mb-16">
                        <h2 class="text-3xl lg:text-4xl font-bold text-gray-900 mb-4">
                            Your Complete Diabetes Management Companion
                        </h2>
                    </div>
                    <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
                        <div class="bg-white p-8 rounded-xl shadow-lg">
                            <div class="w-16 h-16 bg-blue-600 rounded-lg flex items-center justify-center mb-6">
                                <span class="text-white text-2xl">📚</span>
                            </div>
                            <h3 class="text-xl font-semibold text-gray-900 mb-4">Expert-Backed Learning</h3>
                            <p class="text-gray-600">Access clinically-reviewed diabetes education and resources.</p>
                        </div>
                        <div class="bg-white p-8 rounded-xl shadow-lg">
                            <div class="w-16 h-16 bg-green-600 rounded-lg flex items-center justify-center mb-6">
                                <span class="text-white text-2xl">👥</span>
                            </div>
                            <h3 class="text-xl font-semibold text-gray-900 mb-4">Never Walk Alone</h3>
                            <p class="text-gray-600">Join supportive community discussions with others on similar journeys.</p>
                        </div>
                        <div class="bg-white p-8 rounded-xl shadow-lg">
                            <div class="w-16 h-16 bg-purple-600 rounded-lg flex items-center justify-center mb-6">
                                <span class="text-white text-2xl">📊</span>
                            </div>
                            <h3 class="text-xl font-semibold text-gray-900 mb-4">Practical Daily Guidance</h3>
                            <p class="text-gray-600">Get personalized recommendations for meals, exercise, and lifestyle.</p>
                        </div>
                    </div>
                </div>
            </section>

            <section id="community" class="bg-blue-600 py-20">
                <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
                    <h2 class="text-3xl lg:text-4xl font-bold text-white mb-6">
                        Your Success Story Starts Here
                    </h2>
                    <p class="text-xl text-blue-100 mb-8">
                        Join a community achieving diabetes management goals through shared experiences.
                    </p>
                    <a href="/signup" class="bg-white text-blue-600 px-8 py-4 rounded-lg text-lg font-semibold hover:bg-gray-100">
                        Start Your Journey Today
                    </a>
                </div>
            </section>
        `;
}

export default Home;