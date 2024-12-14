package api

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type BlogPost struct {
	Title   string
	Author  string
	Date    string
	Content string
}

type Post struct {
	ID      string
	Title   string
	Excerpt template.HTML
	Author  string
	Date    string
	Content template.HTML
}

type Issue struct {
	StatusCode int
	Problem    string
}

// Initialize variable to hold error message and status codes

var hitch Issue

var LoadTemplate = func() (*template.Template, error) {
	return template.ParseFiles("../frontend/public/error.html")
}

func GlucoseTrackerEndPointHandler(w http.ResponseWriter, r *http.Request) {
	// Capture glucose level and date from the request query parameters
	glucoseLevel := r.URL.Query().Get("glucose")
	glucoseDate := r.URL.Query().Get("date")

	glucoseParam := map[string]string{glucoseLevel: glucoseDate}

	// Set response header and JSON encode the glucose level and date
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(glucoseParam)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]
	posts := map[string]Post{
		"1": {
			Title:  "Understanding Diabetes: A Comprehensive Guide",
			Author: "Dr. Samantha Johnson",
			Date:   "March 15, 2024",
			Content: template.HTML(`<header>
			<h1>Understanding Diabetes: A Comprehensive Guide</h1>
		</header>
	
		<section>
			<h2>Introduction</h2>
			<p>Diabetes is a condition that affects millions of people worldwide, yet many are still unfamiliar with it. Whether you're newly diagnosed or looking to learn more, this guide will introduce you to the basics of diabetes, its types, symptoms, causes, and risk factors. By understanding these key elements, you'll be better equipped to manage the condition, either for yourself or a loved one.</p>
		</section>
	
		<section>
			<h2>What is Diabetes?</h2>
			<p>Diabetes occurs when your body either cannot produce enough insulin or cannot use it effectively. Insulin is a hormone that helps your cells absorb glucose (a type of sugar) from your bloodstream for energy. Without effective insulin action, glucose builds up in your blood, leading to symptoms and potential complications.</p>
		</section>
	
		<section>
			<h2>The Three Main Types of Diabetes</h2>
			<article>
				<h3>Type 1 Diabetes</h3>
				<p><strong>What is it?</strong> An autoimmune condition where the body attacks its insulin-producing cells in the pancreas. It’s often diagnosed in children or young adults.</p>
				<p><strong>How is it managed?</strong> People with Type 1 diabetes must take insulin injections daily to regulate blood sugar.</p>
			</article>
	
			<article>
				<h3>Type 2 Diabetes</h3>
				<p><strong>What is it?</strong> The most common form of diabetes, where the body becomes resistant to insulin or doesn’t produce enough of it. It typically develops in adults, but is increasingly seen in younger individuals due to lifestyle factors.</p>
				<p><strong>How is it managed?</strong> Often managed with lifestyle changes like diet and exercise, though medication or insulin may be required.</p>
			</article>
	
			<article>
				<h3>Gestational Diabetes</h3>
				<p><strong>What is it?</strong> Occurs during pregnancy when the body cannot produce enough insulin. This type usually disappears after childbirth but increases the risk of developing Type 2 diabetes later in life.</p>
				<p><strong>How is it managed?</strong> Managed through diet, exercise, and sometimes insulin, with close monitoring during pregnancy.</p>
			</article>
		</section>
	
		<section>
			<h2>Symptoms of Diabetes</h2>
			<p>Recognizing the symptoms early can help prevent complications. Common signs include:</p>
			<ul>
				<li><strong>Frequent urination:</strong> High blood sugar levels lead to increased urination as the kidneys filter excess glucose.</li>
				<li><strong>Increased thirst:</strong> Frequent urination causes dehydration, which triggers thirst.</li>
				<li><strong>Fatigue:</strong> When glucose isn't used effectively, your cells don't get the energy they need.</li>
				<li><strong>Blurred vision:</strong> High blood sugar can cause fluid to shift in the eyes, leading to blurry vision.</li>
				<li><strong>Unexplained weight loss:</strong> Despite eating well, the body starts breaking down fat and muscle for energy.</li>
			</ul>
			<p>If you experience these symptoms, consult with a healthcare provider for testing and diagnosis.</p>
		</section>
	
		<section>
			<h2>What Causes Diabetes?</h2>
			<p>The causes of diabetes vary depending on the type, but common contributing factors include:</p>
			<ul>
				<li><strong>Genetics:</strong> Family history plays a significant role, particularly for Type 1 and Type 2 diabetes.</li>
				<li><strong>Lifestyle:</strong> Poor diet, lack of exercise, and obesity are major contributors to Type 2 diabetes.</li>
				<li><strong>Autoimmune response:</strong> In Type 1 diabetes, the immune system mistakenly attacks the insulin-producing cells in the pancreas.</li>
				<li><strong>Hormonal changes:</strong> Pregnancy-related hormonal shifts can cause gestational diabetes.</li>
			</ul>
		</section>
	
		<section>
			<h2>Risk Factors for Diabetes</h2>
			<p>Certain factors increase your chances of developing diabetes, including:</p>
			<ul>
				<li><strong>Family history:</strong> Having a close relative with diabetes increases your risk.</li>
				<li><strong>Age:</strong> Risk increases with age, especially after 45, but Type 2 diabetes is becoming more common in younger people.</li>
				<li><strong>Obesity:</strong> Extra fat, especially around the abdomen, raises the risk of insulin resistance.</li>
				<li><strong>Physical inactivity:</strong> A sedentary lifestyle contributes to weight gain and increases the risk of developing Type 2 diabetes.</li>
				<li><strong>Unhealthy diet:</strong> Diets high in sugar, fats, and low in fiber are linked to the development of Type 2 diabetes.</li>
			</ul>
			<p>For women, gestational diabetes can increase the risk of Type 2 diabetes later in life.</p>
		</section>
	
		<section>
			<h2>Preventing and Managing Diabetes</h2>
			<p>While Type 1 diabetes cannot be prevented, Type 2 and gestational diabetes can often be managed or even prevented with lifestyle changes. Here’s how:</p>
			<ul>
				<li><strong>Healthy eating:</strong> A balanced diet with plenty of vegetables, fruits, whole grains, and lean proteins can help control blood sugar.</li>
				<li><strong>Exercise:</strong> Physical activity improves insulin sensitivity and helps manage weight.</li>
				<li><strong>Weight management:</strong> Losing excess weight can prevent or help manage Type 2 diabetes.</li>
				<li><strong>Regular check-ups:</strong> Monitoring blood sugar levels and having routine doctor visits are essential for early detection and long-term management.</li>
			</ul>
		</section>
	
		<footer>
			<h2>Conclusion</h2>
			<p>Diabetes is a serious condition, but with proper knowledge and management, those living with diabetes can lead healthy lives. Whether you're managing the condition yourself or supporting a loved one, understanding the basics is the first step toward effective diabetes care.</p>
			<p>Stay tuned for future blog posts where we’ll dive deeper into managing blood sugar, healthy eating, exercise, and more. With the right care, diabetes is manageable, and a healthy future is within reach.</p>
			<p><strong>Call to Action:</strong></p>
			<p>Have you or someone you know been recently diagnosed with diabetes? What questions do you have about managing the condition? Drop them in the comments below and let’s start a conversation!</p>
			<p>Share this guide with your friends and family to help spread awareness about diabetes!</p>
		</footer>`),
		},
		"2": {
			Title:  "How to Monitor Blood Sugar Levels: A Step-by-Step Guide",
			Author: "Jane Smith",
			Date:   "January 15, 2024",
			Content: template.HTML(`<h1>How to Monitor Blood Sugar Levels: A Step-by-Step Guide</h1>

			<p><strong>Introduction:</strong></p>
			<p>Monitoring blood sugar levels is a crucial part of managing diabetes. Whether you have diabetes or are caring for someone with the condition, understanding how and when to check blood sugar can make a significant difference in overall health and well-being. In this step-by-step guide, we’ll explain how to effectively monitor blood sugar levels, the tools you need, and why regular checks are essential.</p>
		
			<h2>Why Monitoring Blood Sugar is Important:</h2>
			<p>Keeping blood sugar levels within a healthy range is key to preventing complications of diabetes, such as heart disease, nerve damage, and vision problems. By regularly monitoring blood sugar, you can:</p>
			<ul>
				<li>Track how food, exercise, stress, and medication affect blood sugar levels.</li>
				<li>Make informed decisions about diet, exercise, and medication adjustments.</li>
				<li>Identify patterns that help prevent dangerous blood sugar highs and lows.</li>
			</ul>
			<p>Understanding your readings empowers you to manage diabetes more effectively, whether it’s Type 1, Type 2, or gestational diabetes.</p>
		
			<h2>Step 1: Choosing the Right Tool - The Glucometer</h2>
			<p>A glucometer, or blood glucose meter, is the most common tool for checking blood sugar levels. These devices come in different models, but they all work by measuring the glucose level in a small drop of blood, typically obtained from a fingertip.</p>
			<p>Here’s how to choose a glucometer:</p>
			<ul>
				<li><strong>Accuracy:</strong> Look for a glucometer that has been approved by the FDA for accurate readings.</li>
				<li><strong>Ease of Use:</strong> Some meters have larger displays, smaller sample sizes, or Bluetooth connectivity for easier data tracking.</li>
				<li><strong>Cost:</strong> Consider the cost of the glucometer as well as the test strips, which are needed for each measurement.</li>
			</ul>
		
			<h2>Step 2: When to Check Blood Sugar</h2>
			<p>How often you monitor your blood sugar depends on your individual needs and your doctor’s advice. However, there are common times when you should check:</p>
			<ul>
				<li><strong>Before meals:</strong> This helps you understand your baseline glucose level before eating and can guide meal choices.</li>
				<li><strong>After meals:</strong> This shows how your body responds to food and whether blood sugar spikes occur after eating.</li>
				<li><strong>Before bedtime:</strong> A pre-bedtime check can prevent overnight blood sugar lows or highs.</li>
				<li><strong>When you feel off:</strong> If you experience symptoms like dizziness, weakness, or confusion, checking your blood sugar is a good idea to rule out hypo- or hyperglycemia.</li>
			</ul>
		
			<h2>Step 3: How to Check Your Blood Sugar</h2>
			<p>Now that you have your glucometer, here’s a simple, step-by-step guide for checking your blood sugar:</p>
			<ol>
				<li><strong>Wash your hands:</strong> This prevents any dirt or residue from affecting your test results.</li>
				<li><strong>Prepare the glucometer:</strong> Insert a test strip into the glucometer and prepare your lancing device (the tool used to prick your finger).</li>
				<li><strong>Prick your finger:</strong> Use the lancing device to prick the side of your fingertip. This should be quick and relatively painless.</li>
				<li><strong>Place the blood sample on the test strip:</strong> Gently squeeze your finger to get a small drop of blood. Place it on the test strip in the glucometer.</li>
				<li><strong>Read the result:</strong> Your glucometer will show your blood sugar level within seconds. Make a note of the result for tracking purposes.</li>
			</ol>
		
			<h2>Step 4: Interpreting Your Results</h2>
			<p>Understanding your readings is crucial for effective management. Generally, blood sugar levels are categorized as follows:</p>
			<ul>
				<li><strong>Before meals:</strong> 70-130 mg/dL (milligrams per deciliter) is considered normal.</li>
				<li><strong>After meals (2 hours):</strong> Less than 180 mg/dL is ideal for most people with diabetes.</li>
			</ul>
			<p>If your readings are consistently outside these ranges, talk to your healthcare provider for guidance.</p>
		
			<h2>Step 5: Keeping Track of Your Readings</h2>
			<p>Keeping a record of your blood sugar levels helps you and your healthcare provider make informed decisions. You can use a notebook or digital tools like apps connected to your glucometer for tracking trends.</p>
		
			<h2>Conclusion:</h2>
			<p>Regular blood sugar monitoring is vital to managing diabetes and preventing complications. With the right tools and knowledge, you can take control of your health. Remember, consistency is key—stick to a routine, follow your doctor’s advice, and track your readings regularly. Monitoring blood sugar can help you make the necessary adjustments to your lifestyle, diet, and treatment plan to stay healthy and active.</p>`),
		},
		"3": {
			Title:  "How to Recognize and Prevent Diabetes Complications",
			Author: "Michael Johnson",
			Date:   "December 20, 2023",
			Content: template.HTML(`<p><strong>"How to Recognize and Prevent Diabetes Complications"</strong></p>

			<p><span>Diabetes is a chronic condition that, if not well-managed, can lead to a range of serious health complications. Recognizing the early signs of complications and taking preventive measures can make a huge difference in maintaining quality of life for people with diabetes. In this post, we’ll cover common complications and practical steps you can take to prevent them.</span></p>
		
			<h2><span>Common Diabetes Complications</span></h2>
			<p>Diabetes can lead to complications that affect various parts of the body. The most common complications include:</p>
			<ul>
				<li><strong>Heart Disease:</strong> Diabetes increases the risk of heart disease due to high blood sugar levels damaging blood vessels.</li>
				<li><strong>Neuropathy:</strong> High blood sugar can damage the nerves, leading to numbness, tingling, or pain, usually in the feet and hands.</li>
				<li><strong>Kidney Disease:</strong> Diabetes can damage the kidneys over time, leading to kidney failure in severe cases.</li>
				<li><strong>Eye Problems:</strong> Diabetic retinopathy can cause damage to the blood vessels in the eyes, leading to vision loss if untreated.</li>
			</ul>
		
			<h2><span>How to Prevent Diabetes Complications</span></h2>
			<p>Preventing complications begins with effective management of blood sugar levels. Here’s what you can do:</p>
			<ul>
				<li><strong>Monitor Blood Sugar Regularly:</strong> Frequent monitoring helps keep track of your blood glucose levels, which is crucial for preventing complications.</li>
				<li><strong>Adopt a Healthy Diet:</strong> A balanced diet low in processed foods, sugars, and fats helps control blood sugar levels and promotes heart health.</li>
				<li><strong>Exercise Regularly:</strong> Physical activity helps regulate blood sugar, reduce stress, and improve circulation.</li>
				<li><strong>Manage Stress:</strong> Chronic stress can negatively impact blood sugar levels. Finding ways to relax is key.</li>
				<li><strong>Get Regular Check-ups:</strong> Routine medical visits help catch early signs of complications, allowing for early intervention.</li>
			</ul>
		
			`),
		},
		"4": {
			Title:  "The Role of Exercise in Managing Diabetes",
			Author: "Michael Johnson",
			Date:   "October 20, 2024",
			Content: template.HTML(`<h1>The Role of Exercise in Managing Diabetes</h1>
			
			<p><strong>Introduction:</strong> Physical activity plays a crucial role in managing diabetes. Regular exercise helps control blood sugar levels, improves cardiovascular health, and boosts overall well-being. Whether you have Type 1, Type 2, or gestational diabetes, incorporating physical activity into your daily routine can make a significant impact on managing the condition. In this post, we’ll explore the benefits of exercise for people with diabetes, the best types of exercises, and tips for incorporating them into your daily life.</p>
		
			<h2><span>Why Exercise Matters for Diabetes Management</span></h2>
			<p>Exercise is an effective way to manage blood sugar levels and improve insulin sensitivity. For people with diabetes, physical activity has numerous benefits, including:</p>
			<ul>
				<li><strong>Improved Blood Sugar Control:</strong> Exercise helps muscles absorb glucose, lowering blood sugar levels. Regular physical activity can also reduce the need for medication in some cases.</li>
				<li><strong>Weight Management:</strong> Physical activity helps burn calories and maintain a healthy weight, which is essential for managing Type 2 diabetes.</li>
				<li><strong>Reduced Risk of Heart Disease:</strong> People with diabetes are at a higher risk for heart disease, but regular exercise strengthens the heart and reduces blood pressure.</li>
				<li><strong>Increased Insulin Sensitivity:</strong> Exercise improves how the body responds to insulin, allowing for better blood sugar control.</li>
			</ul>
		
			<h2><span>The Best Types of Exercise for People with Diabetes</span></h2>
			<p>Not all exercises are created equal when it comes to diabetes management. It’s essential to choose activities that are both effective and sustainable. The best exercises for people with diabetes include:</p>
			
			<h3><span>Aerobic Exercise</span></h3>
			<p>Aerobic activities, such as walking, swimming, cycling, and dancing, are great for improving cardiovascular health and lowering blood sugar levels. These exercises help the body use oxygen more efficiently and improve blood circulation. Aim for at least 150 minutes of moderate aerobic activity per week.</p>
		
			<h3><span>Strength Training</span></h3>
			<p>Strength training involves lifting weights or using resistance bands to build muscle mass. Building muscle helps the body use glucose more effectively and improves insulin sensitivity. It’s recommended to incorporate strength training into your routine two to three times a week.</p>
		
			<h3><span>Flexibility and Balance Exercises</span></h3>
			<p>Stretching and balance exercises, such as yoga or tai chi, can improve flexibility, reduce stress, and enhance overall well-being. These exercises also promote joint health and prevent injuries, which is particularly important for those with neuropathy.</p>
		
			<h2><span>How to Incorporate Exercise into Your Daily Life</span></h2>
			<p>Starting an exercise routine can be challenging, but with the right approach, it can become a manageable and enjoyable part of your daily routine. Here are some tips for incorporating exercise into your life:</p>
			<ul>
				<li><strong>Start Small:</strong> Begin with small goals, such as a 10-minute walk after meals, and gradually increase the duration and intensity.</li>
				<li><strong>Set a Schedule:</strong> Consistency is key to seeing results. Schedule exercise at a time that works best for you, whether it’s in the morning, afternoon, or evening.</li>
				<li><strong>Involve Others:</strong> Exercise is more enjoyable when done with a friend, family member, or support group. Consider joining a walking group or taking a fitness class.</li>
				<li><strong>Monitor Blood Sugar Levels:</strong> Be sure to check your blood sugar before and after exercise to ensure that it stays within a healthy range.</li>
				<li><strong>Listen to Your Body:</strong> It's important to pay attention to how your body responds to exercise. If you feel lightheaded or dizzy, stop and rest.</li>
			</ul>
		
			<h2><span>Conclusion</span></h2>
			<p>Exercise is an essential component of diabetes management. By incorporating regular physical activity into your routine, you can improve your blood sugar control, reduce your risk of complications, and enhance your overall health. Whether it’s aerobic exercise, strength training, or flexibility exercises, find activities that you enjoy and make them a consistent part of your life. As always, consult with your healthcare provider before starting a new exercise regimen to ensure it’s safe and appropriate for your individual needs.</p>`),
		},
		"5": {
			Title:  "Stress and Diabetes: Understanding the Connection",
			Author: "Michael Johnson",
			Date:   "April 20, 2024",
			Content: template.HTML(`<h1>Stress and Diabetes: Understanding the Connection</h1>

			<p><strong>Introduction:</strong> Managing stress is a crucial part of diabetes care. While diabetes requires careful attention to blood sugar levels through diet, exercise, and medication, stress is often an overlooked factor that can affect blood sugar control. Understanding the relationship between stress and diabetes is key to improving overall health and managing the condition more effectively. In this article, we’ll explore how stress impacts blood sugar levels and offer strategies for managing stress to improve diabetes outcomes.</p>
		
			<h2><span>The Link Between Stress and Blood Sugar</span></h2>
			<p>When you experience stress, your body reacts by releasing stress hormones like cortisol and adrenaline. These hormones trigger the “fight-or-flight” response, which prepares your body for immediate action. However, this physiological response can have unintended consequences for individuals with diabetes.</p>
			
			<p><strong>Cortisol</strong>, the primary stress hormone, increases blood sugar levels by stimulating the liver to release more glucose into the bloodstream. In healthy individuals, insulin helps regulate these elevated blood sugar levels. However, in people with diabetes, the body either doesn’t produce enough insulin or doesn’t use it effectively, which can result in higher-than-normal blood sugar levels.</p>
		
			<p>In addition to cortisol, <strong>adrenaline</strong> increases heart rate and blood pressure, preparing your body for action. While the immediate effects may be temporary, prolonged exposure to stress can cause sustained high blood sugar levels, leading to poor diabetes management and increased risk of complications over time.</p>
		
			<h2><span>How Stress Affects Diabetes Outcomes</span></h2>
			<p>Stress can significantly impact diabetes in several ways:</p>
			<ul>
				<li><strong>Increased Blood Sugar Levels:</strong> As mentioned, stress hormones can raise blood sugar levels, making it harder to maintain healthy glucose control.</li>
				<li><strong>Disrupted Routines:</strong> Stress can cause changes in eating habits, sleep patterns, and physical activity, all of which play an important role in blood sugar management.</li>
				<li><strong>Emotional Eating:</strong> Stress can trigger emotional eating, often leading to poor food choices that may spike blood sugar levels.</li>
				<li><strong>Decreased Motivation:</strong> When stressed, individuals may feel less motivated to stick to their diabetes management plan, including exercise, medication adherence, and blood sugar monitoring.</li>
			</ul>
			<iframe 
        	width="560" 
        	height="315" 
        	src="https://www.youtube-nocookie.com/embed/watch?v=jl-eMcz-7Bw" 
        	title="Relationship between Stress and Diabetes" 
			frameborder="0" 
			allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" 
			allowfullscreen>
    		</iframe>
		
			<h2><span>Strategies for Managing Stress</span></h2>
			<p>Managing stress effectively can lead to better diabetes control. Here are some strategies that can help reduce stress and its impact on blood sugar levels:</p>
		
			<h3><span>1. Regular Physical Activity</span></h3>
			<p>Exercise is one of the most effective ways to manage stress. Physical activity boosts the production of endorphins, which are natural stress relievers. Exercise also helps regulate blood sugar levels and improves insulin sensitivity, making it an essential part of diabetes management.</p>
		
			<h3><span>2. Mindfulness and Meditation</span></h3>
			<p>Mindfulness practices, such as meditation and deep breathing exercises, can help calm the mind and reduce stress. Taking a few minutes each day to focus on your breath can lower cortisol levels and improve emotional well-being. These practices can also help with emotional eating and support healthier decision-making regarding diet and lifestyle.</p>
		
			<h3><span>3. Get Enough Sleep</span></h3>
			<p>Lack of sleep can elevate stress and disrupt blood sugar regulation. Aim for 7-9 hours of sleep each night to give your body time to recover and manage stress effectively. Good sleep hygiene, such as sticking to a regular sleep schedule and creating a relaxing bedtime routine, can improve both stress and blood sugar control.</p>
		
			<h3><span>4. Support Network</span></h3>
			<p>Talking to friends, family, or a counselor can help alleviate stress. Having a strong support system can provide emotional comfort and practical advice for managing diabetes. Support groups, either in person or online, can also provide a sense of community and shared experience.</p>
		
			<h2><span>Conclusion</span></h2>
			<p>Stress is an inevitable part of life, but understanding its impact on diabetes can empower you to take control of your health. By managing stress through regular exercise, mindfulness, adequate sleep, and a strong support system, you can reduce its negative effects on blood sugar levels and improve your diabetes management. Remember, addressing stress is just as important as managing diet and medication in maintaining healthy blood sugar levels and preventing complications.</p>`),
		},
	}

	post, ok := posts[postID]
	if !ok {
		NotFoundHandler(w)
		return
	}

	tmpl, err := template.ParseFiles(
		"../frontend/public/base.html",
		"../frontend/public/blog_display.html",
	)
	if err != nil {
		InternalServerErrorHandler(w)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", post)
	if err != nil {
		InternalServerErrorHandler(w)
		return
	}
}

func BlogHomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"../frontend/public/base.html",
		"../frontend/public/blog_home.html",
	)

	data := struct {
		Title string
		Posts []Post
	}{
		Title: "Home",
		Posts: []Post{
			{ID: "1", Title: "Understanding Diabetes: A Comprehensive Guide", Excerpt: template.HTML(`Diabetes occurs when your body either cannot produce enough insulin or cannot use it effectively. Insulin is a hormone that helps your cells absorb glucose (a type of sugar) from your bloodstream for energy. Without effective insulin action, glucose builds up in your blood, leading to symptoms and potential complications.</span></p>
			<p><br></p>`)},
			{ID: "2", Title: "How to Monitor Blood Sugar Levels: A Step-by-Step Guide", Excerpt: template.HTML(`<p>Monitoring blood sugar levels is a crucial part of managing diabetes. Whether you have diabetes or are caring for someone with the condition, understanding how and when to check blood sugar can make a significant difference in overall health and well-being. In this step-by-step guide, we’ll explain how to effectively monitor blood sugar levels, the tools you need, and why regular checks are essential.</p>`)},
			{ID: "3", Title: "How to Recognize and Prevent Diabetes Complications", Excerpt: template.HTML(`<p><span>Diabetes is a chronic condition that, if not well-managed, can lead to a range of serious health complications. Recognizing the early signs of complications and taking preventive measures can make a huge difference in maintaining quality of life for people with diabetes. In this post, we’ll cover common complications and practical steps you can take to prevent them.</span></p>`)},
			{ID: "4", Title: "The Role of Exercise in Managing Diabetes", Excerpt: template.HTML(`<p>Physical activity plays a crucial role in managing diabetes. Regular exercise helps control blood sugar levels, improves cardiovascular health, and boosts overall well-being. Whether you have Type 1, Type 2, or gestational diabetes, incorporating physical activity into your daily routine can make a significant impact on managing the condition. In this post, we’ll explore the benefits of exercise for people with diabetes, the best types of exercises, and tips for incorporating them into your daily life.</p>`)},
			{ID: "5", Title: "Stress and Diabetes: Understanding the Connection", Excerpt: template.HTML(`Managing stress is a crucial part of diabetes care. While diabetes requires careful attention to blood sugar levels through diet, exercise, and medication, stress is often an overlooked factor that can affect blood sugar control. Understanding the relationship between stress and diabetes is key to improving overall health and managing the condition more effectively. In this article, we’ll explore how stress impacts blood sugar levels and offer strategies for managing stress to improve diabetes outcomes.</p>`)},
		},
	}

	if err != nil {
		InternalServerErrorHandler(w)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		InternalServerErrorHandler(w)
		return
	}
}

func BadRequestHandler(w http.ResponseWriter) {
	tmpl, err := LoadTemplate()
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}

	hitch.StatusCode = http.StatusBadRequest
	hitch.Problem = "Bad Request!"

	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}

func InternalServerErrorHandler(w http.ResponseWriter) {
	tmpl, err := LoadTemplate()
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}

	hitch.StatusCode = http.StatusInternalServerError
	hitch.Problem = "Internal Server Error!"

	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}

func NotFoundHandler(w http.ResponseWriter) {
	tmpl, err := LoadTemplate()
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}

	hitch.StatusCode = http.StatusNotFound
	hitch.Problem = "Not Found!"

	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}
