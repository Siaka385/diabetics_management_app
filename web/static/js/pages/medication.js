export async function renderMedication() {
    try {
        const response = await fetch('/medication');
        if (!response.ok) throw new Error('Failed to fetch');
        
        const userData = await response.json();
        const app = document.getElementById('app');
        if (!app) return;
        
        app.innerHTML = `
            <div class="min-h-screen bg-gray-50 flex">
                ${renderSidebar(userData, '/medication')}
                <div class="flex-1 p-8">
                    <h1 class="text-3xl font-bold mb-8">Medication Management</h1>
                    
                    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
                        <div class="lg:col-span-2 space-y-6">
                            <div class="bg-white p-6 rounded-xl shadow-sm border">
                                <h2 class="text-xl font-semibold mb-6">Current Medications</h2>
                                <div class="space-y-4">
                                    ${renderMedications()}
                                </div>
                            </div>
                        </div>
                        
                        <div class="space-y-6">
                            <div class="bg-white p-6 rounded-xl shadow-sm border">
                                <button id="addMedicationBtn" class="w-full bg-blue-600 text-white py-3 px-4 rounded-lg hover:bg-blue-700 transition-colors mb-6">
                                    Add Medication
                                </button>
                                
                                <h3 class="text-lg font-semibold mb-4">Upcoming Reminders</h3>
                                <div class="space-y-3">
                                    <div class="flex items-center p-3 bg-blue-50 rounded-lg">
                                        <div class="w-10 h-10 bg-blue-600 rounded-full flex items-center justify-center mr-3">
                                            <span class="text-white">💊</span>
                                        </div>
                                        <div>
                                            <div class="font-medium text-sm">Metformin</div>
                                            <div class="text-xs text-gray-500">Next dose at 6:00 PM</div>
                                        </div>
                                    </div>
                                    <div class="flex items-center p-3 bg-green-50 rounded-lg">
                                        <div class="w-10 h-10 bg-green-600 rounded-full flex items-center justify-center mr-3">
                                            <span class="text-white">💉</span>
                                        </div>
                                        <div>
                                            <div class="font-medium text-sm">Insulin Glargine</div>
                                            <div class="text-xs text-gray-500">Next dose at 9:00 PM</div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            
            <!-- Add Medication Modal -->
            <div id="addMedicationModal" class="fixed inset-0 bg-black bg-opacity-50 hidden items-center justify-center z-50">
                <div class="bg-white p-6 rounded-xl max-w-md w-full mx-4">
                    <div class="flex justify-between items-center mb-4">
                        <h2 class="text-xl font-semibold">Add New Medication</h2>
                        <button id="closeModal" class="text-gray-500 hover:text-gray-700 text-2xl">&times;</button>
                    </div>
                    <form id="addMedicationForm" class="space-y-4">
                        <div>
                            <label class="block mb-2 font-medium text-gray-700">Medication Name</label>
                            <input type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:border-blue-500" placeholder="Enter medication name">
                        </div>
                        <div>
                            <label class="block mb-2 font-medium text-gray-700">Dosage</label>
                            <input type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:border-blue-500" placeholder="e.g., 500mg">
                        </div>
                        <div>
                            <label class="block mb-2 font-medium text-gray-700">Frequency</label>
                            <select required class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:border-blue-500">
                                <option value="">Select Frequency</option>
                                <option value="daily">Once Daily</option>
                                <option value="twice-daily">Twice Daily</option>
                                <option value="three-times-daily">Three Times Daily</option>
                                <option value="as-needed">As Needed</option>
                            </select>
                        </div>
                        <div>
                            <label class="block mb-2 font-medium text-gray-700">Reminder Time</label>
                            <input type="time" required class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:border-blue-500">
                        </div>
                        <button type="submit" class="w-full bg-green-600 text-white py-2 px-4 rounded-md hover:bg-green-700 transition-colors">
                            Add Medication
                        </button>
                    </form>
                </div>
            </div>
        `;
        
        // Add event handlers
        document.getElementById('addMedicationBtn').addEventListener('click', () => {
            document.getElementById('addMedicationModal').classList.remove('hidden');
            document.getElementById('addMedicationModal').classList.add('flex');
        });
        
        document.getElementById('closeModal').addEventListener('click', () => {
            document.getElementById('addMedicationModal').classList.add('hidden');
            document.getElementById('addMedicationModal').classList.remove('flex');
        });
        
        document.getElementById('addMedicationForm').addEventListener('submit', (e) => {
            e.preventDefault();
            alert('Medication added successfully!');
            document.getElementById('addMedicationModal').classList.add('hidden');
            document.getElementById('addMedicationModal').classList.remove('flex');
            e.target.reset();
        });
        
        // Add medication action handlers
        document.querySelectorAll('.btn-taken').forEach(btn => {
            btn.addEventListener('click', () => {
                alert('Medication marked as taken!');
                btn.classList.add('opacity-50');
                btn.disabled = true;
            });
        });
        
        document.querySelectorAll('.btn-skip').forEach(btn => {
            btn.addEventListener('click', () => {
                if (confirm('Are you sure you want to skip this medication?')) {
                    alert('Medication skipped. Please consult your doctor if you regularly miss doses.');
                }
            });
        });
        
    } catch (error) {
        window.location.href = '/login';
    }
}

function renderMedications() {
    const medications = [
        {
            name: "Metformin",
            dosage: "500mg",
            frequency: "2x daily",
            nextDose: "6:00 PM"
        },
        {
            name: "Insulin Glargine", 
            dosage: "10 units",
            frequency: "Once daily",
            nextDose: "9:00 PM"
        }
    ];
    
    return medications.map(med => `
        <div class="flex justify-between items-center p-4 bg-gray-50 rounded-lg">
            <div class="flex-grow">
                <div class="font-semibold text-blue-600">${med.name}</div>
                <div class="text-sm text-gray-600">${med.dosage} - ${med.frequency}</div>
                <div class="text-xs text-gray-500">Next dose: ${med.nextDose}</div>
            </div>
            <div class="flex gap-2">
                <button class="btn-taken bg-green-600 text-white px-3 py-1 rounded text-sm hover:bg-green-700 transition-colors">
                    Taken
                </button>
                <button class="btn-skip bg-red-600 text-white px-3 py-1 rounded text-sm hover:bg-red-700 transition-colors">
                    Skip
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