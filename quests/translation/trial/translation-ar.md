---
title: "النشر المحكوم بالوكلاء: لماذا تقفل الآلة قرار النجاح"
date: 2026-06-10
lastmod: 2026-06-11
tags: [agents, publishing, quality-gates]
---

![آلية سقاطة توضح التقدم أحادي الاتجاه](/images/ratchet-gate.png)

عندما يكتب وكيل ذكي تدوينة وينشرها من البداية إلى النهاية، فإن السؤال الصعب ليس
جودة التوليد بل سلطة التحقق. التوليد احتمالي، أما التحقق فيجب أن يكون حتمياً،
ولا ينبغي السماح إلا للجانب الحتمي بإعلان اكتمال المهمة.

## نموذج السقاطة

الكويست مجموعة من العناصر تُقفل بسقاطة أحادية الاتجاه. متى اجتاز عنصر البوابة
فلن يتراجع بصمت أبداً: العمل المتبقي لا يفعل سوى أن يتقلص. الوكيل يقترح والبوابة
تقرر — انظر بوابة [Robots Exclusion Protocol](https://www.rfc-editor.org/rfc/rfc9309.html)
من جولة سابقة لترى كيف يرفض فحص الاستشهاد المصادر غير القابلة للوصول.

أبسط قاعدة بوابة تبدو هكذا:

```go
var ruleExample = gate.Rule{
	Meta: gate.RuleMeta{ID: "example", Level: gate.LevelFail},
	Check: func(ctx gate.Context) (bool, quest.Fact) {
		// fired == a violation was found

		return false, quest.Fact{}
	},
}
```

## الترجمة بوصفها كويست

الترجمات عناصر أيضاً: مقالة واحدة في N لغة. تُحكم كل لغة على حدة، فإخفاق تسليم
بالعربية لا يعطل أبداً نجاح تسليم باليابانية. العقد البنيوي — العناوين وكتل
الفقرات وأسوار الشيفرة والروابط — يُقارن بالأصل، كما شرحنا في
[المقالة الأولى](/ar/posts/agent-gated-publishing/) من هذه السلسلة. وضع Hugo
متعدد اللغات موثق في
[Hugo multilingual guide](https://gohugo.io/content-management/multilingual/).

## المصادر

- [RFC 9309: Robots Exclusion Protocol](https://www.rfc-editor.org/rfc/rfc9309.html)
- [Hugo multilingual guide](https://gohugo.io/content-management/multilingual/)
