package tests

import (
	"os"
	"testing"
)

func split(str string, lim int) []string {
	buf := []rune(str)
	var chunk []rune
	chunks := make([]string, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, string(chunk))
	}
	if len(buf) > 0 {
		chunks = append(chunks, string(buf))
	}
	return []string(chunks)
}

func TestSplit(t *testing.T) {
	// String, which is tested
	const testStr = "Во всех временах дружество почитали из числа первых благ в жизни; " +
		"сие чувствование родится вместе с нами; первое движение сердца состоит в том, чтобы искать соединиться с " +
		"другим сердцем, и между тем целый свет жалуется, что нет друзей. С начала мира все веки вместе едва-едва " +
		"произвели три или четыре примера дружества совершенного. Но если все люди согласны, что дружество " +
		"прелестно, почто же не ищут наслаждаться сим благом? Не есть ли сие заблуждение слепого человечества и " +
		"следствие развращения оного - желать блаженства, иметь его в своих руках и убегать его? " +
		"Выгоды дружества блистательны сами собою: вся природа единогласно подтверждает, что они приятнейшие изо " +
		"всех благ земных: без дружества жизнь теряет свои приятности, человек, оставленный самому себе, чувствует " +
		"в своем сердце пустоту, которую единое дружество наполнить может; от природы заботливый и беспокойный, " +
		"в недрах дружества утишает он свои чувствования. Коль полезно пристанище дружбы! Она охраняет от " +
		"коварства людей, которые почти все непостоянны, обманчивы и лживы. Первое достоинство дружбы есть " +
		"вспомоществовать добрым советом. Сколь бы ни рассудителен кто был, но всегда нужен проводник; не должно " +
		"без опасения вверяться своему собственному разуму, который страсти наши заставляют часто говорить по их " +
		"воле. Древние познали все благо любви, но они описания дружества сделали столь огромными, что заставили " +
		"почитать оное за прекрасную выдумку, которой нет в природе. Кажется, они худо знали свойства человека, " +
		"когда умышляли прельщать его такими описаниями и заставлять искать дружбы, столь богато раскрашенной ими: " +
		"они как будто позабыли, что человек более склонен знатным примерам удивляться, нежели им последовать. Но " +
		"более ли ныне умели познать дружество? К стыду нашего века, кажется, признают, что делить с другом своим " +
		"имение свое есть вышняя степень дружества и величайшее ее усилие. Но люди, которые так рассуждают, " +
		"не изъявляют ли сим одно развращение нынешнего времени, алчность и привязанность свою к богатству, " +
		"и способны ли они чувствовать дружество? В таких описаниях, желая дать почувствовать цену дружества, " +
		"не оное, но свое корыстолюбие они изображают: главнейшая выгода дружества состоит в том, чтоб в друге " +
		"своем найти образец добродетелей; нам всегда лестно приобретать почтение любимой особы, и сие-то желание " +
		"заставляет нас последовать добродетелям, которым мы в ней удивляемся. Сенека советует другу своему выбрать " +
		"почтеннейшего между людьми, поступать всегда, как будто б был в его присутствии, и давать ему отчет в своих " +
		"делах: сей великий человек, сей судья наших деяний есть наш друг; никто столь за нас не отвечает нам " +
		"самим, никто не оправдывает нас столь в глазах общества, как друг почтенный. Нам непростительно быть " +
		"несовершенными в его глазах; и потому никогда не видно порочного в дружестве с добродетельным. " +
		"Неприятно видеть того, кто судит и обвиняет нас беспрестанно. Надобно быть прежде уверенными хорошо о " +
		"беспорочности своего сердца и правил, если хочешь выбрать себе друга. Выбирая друга, ты судью себе " +
		"избираешь: порочный удобен ли отважиться на такой выбор и выбор сей не принадлежит ли единой добродетели? " +
		"Я боюсь, говорил Плиний, лишась своего друга, я боюсь, чтобы не ослабеть в пути добродетели; уже потерял " +
		"я моего проводника и свидетеля моей жизни. Дружество по необходимости заставляет нас быть добродетельными: " +
		"оно не может храниться иначе, как между почтенными людьми, и тот, кто желает сохранить дружество, " +
		"должен стараться подражать в добродетели своему другу. Итак, в дружбе находятся и полезные советы, " +
		"возбуждения последовать добрым примерам; друг с нами делит наши печали, помогает нам в наших нуждах, " +
		"не побужденный ни просьбою, ни корыстию, ни исканием. Рассмотрим же теперь истинное свойство дружества, " +
		"по коему бы можно было познавать оное. Первое достоинство, которое должно сыскивать в друге, есть " +
		"добродетель: она-то уверяет нас в нем, что он способен к дружеству и оного достоин. Не надейся нимало " +
		"на ваши обязательства, как скоро не на сем основании они утверждены: ныне не выбор, но нужды соединяют " +
		"людей, и для того-то нынешнее дружество так же скоро кончится, как и начинается: дружатся без разбору и " +
		"ссорятся не раздумывая; ничто столь не презренно: худой выбор оказывает или дурное сердце, или дурной " +
		"разум. Из тысячи умей выбрать себе друга, ничто столь не важно, как сей выбор, ибо от него зависит наше " +
		"благополучие. Помысли, что нас судят в наших друзьях; выбрать друга - значит дать обществу свой образ и " +
		"открыть свое сердце. Всякий бы трепетал, если бы размыслил, сколь важно признать кого своим другом. Хочешь " +
		"ли быть почтен? Обращайся с людьми почтенными, итак, надобно хорошо познать прежде, нежели войти в " +
		"обязательствы. Первый знак, могущий уверить, что выбранный достоин дружества, есть добродетель. Потом " +
		"должно, чтобы выбираемый друг был свободен и неподверженный страстям: одержимые гордостию и надмением не " +
		"способны к таким чувствованиям, а еще менее те, кои окованы узами любви. Любовь похищает всю цену " +
		"дружества; она есть страсть жестокая, а дружба - ощущение приятное и спокойное. Любовь напояет душу " +
		"некоим восторгом радости, за коею нередко следует жестокое сокрушение; но дружба есть наслаждение рассудка " +
		"всегда чистое и всегда ровное; ничто не может его переменить, ни ослабить; оно питает душу, и время не " +
		"сокрушает, но подкрепляет его. Занимающиеся исканием почести также не способны к сему чувствованию: " +
		"они более ищут богатства, нежели дружества, в дружестве не менее нужны чистые нравы; тот подвергается " +
		"великой опасности, кто вступает в дружбу с человеком развращенного поведения. Я думаю также, что чрезмерная " +
		"молодость не способна к наслаждению совершенным дружеством: мы видим много молодых людей, которые " +
		"называются друзьями и думают, что это истина, но их связывают одни забавы; а забавы не суть узы, достойные " +
		"дружества. Мы в летах, приличных дружеству, говорит Сенека своему другу: сильнейшие страсти в тебе угасли, " +
		"остались одни приятные, и мы будем наслаждаться забавами дружества. Но рассмотрим должности дружбы. " +
		"Три время находятся в дружбе: начало, продолжение и конец. Как всякое начало дружества наполнено " +
		"чувствованиями, то все наполнено забавою при рождающемся дружестве. Но часто случается, что вкус " +
		"притупляется привычкою, очарование исчезает, и тогда уже должно поддерживать дружество по рассуждению; " +
		"подпора всегда слабая. В дружбе, как в любви, должно щадить чувства и вкус от притупления: сия " +
		"предосторожность позволена. Но кто может воздержаться от утешения невинного и позволенного? Итак, " +
		"должно утверждать дружество на крепчайшем основании; почтение, подкрепляемое знанием достоинств; " +
		"вот основание, которое никогда не может поколебаться. Любовь пишется с завязанными глазами; но у " +
		"дружества снимается сия завязка; дружество рассматривает прежде, нежели избирает, и привязывается к " +
		"одним достоинствам особы: ибо те одни достойны быть любимы, которые в себе самих заключают причину, " +
		"побуждающую их любить. Сделав добрый выбор, надобно на нем утвердиться и иметь к своему другу непременчивое " +
		"почтение и основанное не на одних чувствах, но на рассуждении, ибо чувства могут охладеть, и тогда " +
		"почтение, сопряженное со справедливостию, только остается непременным. Не надобно дозволять себе ценить " +
		"слабости в своем друге и еще менее говорить о них; должно почитать дружбу. Но как она дана нам к тому, " +
		"чтоб быть помошию добродетели, а не участвовать в пороках, то должно уведомлять, когда друг твой " +
		"уклонится с пути добродетели; если он упрямствует, вооружись властию и силою, которую подают премудрые " +
		"советы и чистота добрых намерений. Не думай, что по окончании дружбы ты уже ничем не обязан: " +
		"обязательства труднейшие остаются тебе исполнить, и в коих одна честь может тебя поддержать; к старому " +
		"дружеству должно хранить почтение. Некоторые думают, что смерть разрушает все обязательства, и малое " +
		"число людей умеют быть друзьями умерших, хотя великолепнейшее украшение на наших похоронах суть слезы " +
		"наших друзей и приятнейшая нам гробница - их сердца; не думай, однакож, чтоб слезы, проливаемые только " +
		"на погребении твоего друга, были последним знаком твоей к нему горячности: ты еще остаешься обязанным " +
		"их имени, славе и дому, он должен жить в твоем сердце, в твоем воспоминании, в твоих устах похвалами и " +
		"в твоих поступках подражанием его добродетелям. Вот правила, вот образ дружества, начертанный на моем " +
		"сердце самою природою, и я тогда только почту себя несчастливым, когда сердце мое охладеет к толь " +
		"благородным чувствованиям; когда дружество меня оставит и свет претворится для моего сердца в пустыню."

	// The output that telegram returns
	var testArr = []string{"Во всех временах дружество почитали из числа первых благ в жизни; сие чувствование " +
		"родится вместе с нами; первое движение сердца состоит в том, чтобы искать соединиться с другим сердцем, " +
		"и между тем целый свет жалуется, что нет друзей. С начала мира все веки вместе едва-едва произвели три " +
		"или четыре примера дружества совершенного. Но если все люди согласны, что дружество прелестно, почто же " +
		"не ищут наслаждаться сим благом? Не есть ли сие заблуждение слепого человечества и следствие развращения " +
		"оного - желать блаженства, иметь его в своих руках и убегать его? Выгоды дружества блистательны сами " +
		"собою: вся природа единогласно подтверждает, что они приятнейшие изо всех благ земных: без дружества жизнь " +
		"теряет свои приятности, человек, оставленный самому себе, чувствует в своем сердце пустоту, которую единое " +
		"дружество наполнить может; от природы заботливый и беспокойный, в недрах дружества утишает он свои " +
		"чувствования. Коль полезно пристанище дружбы! Она охраняет от коварства людей, которые почти все " +
		"непостоянны, обманчивы и лживы. Первое достоинство дружбы есть вспомоществовать добрым советом. Сколь " +
		"бы ни рассудителен кто был, но всегда нужен проводник; не должно без опасения вверяться своему " +
		"собственному разуму, который страсти наши заставляют часто говорить по их воле. Древние познали все " +
		"благо любви, но они описания дружества сделали столь огромными, что заставили почитать оное за " +
		"прекрасную выдумку, которой нет в природе. Кажется, они худо знали свойства человека, когда умышляли " +
		"прельщать его такими описаниями и заставлять искать дружбы, столь богато раскрашенной ими: они как " +
		"будто позабыли, что человек более склонен знатным примерам удивляться, нежели им последовать. Но более " +
		"ли ныне умели познать дружество? К стыду нашего века, кажется, признают, что делить с другом своим имение " +
		"свое есть вышняя степень дружества и величайшее ее усилие. Но люди, которые так рассуждают, не изъявляют " +
		"ли сим одно развращение нынешнего времени, алчность и привязанность свою к богатству, и способны ли они " +
		"чувствовать дружество? В таких описаниях, желая дать почувствовать цену дружества, не оное, но свое " +
		"корыстолюбие они изображают: главнейшая выгода дружества состоит в том, чтоб в друге своем найти образец " +
		"добродетелей; нам всегда лестно приобретать почтение любимой особы, и сие-то желание заставляет нас " +
		"последовать добродетелям, которым мы в ней удивляемся. Сенека советует другу своему выбрать " +
		"почтеннейшего между людьми, поступать всегда, как будто б был в его присутствии, и давать ему отчет " +
		"в своих делах: сей великий человек, сей судья наших деяний есть наш друг; никто столь за нас не отвечает " +
		"нам самим, никто не оправдывает нас столь в глазах общества, как друг почтенный. Нам непростительно быть " +
		"несовершенными в его глазах; и потому никогда не видно порочного в дружестве с добродетельным. Неприятно " +
		"видеть того, кто судит и обвиняет нас беспрестанно. Надобно быть прежде уверенными хорошо о беспорочности " +
		"своего сердца и правил, если хочешь выбрать себе друга. Выбирая друга, ты судью себе избираешь: порочный " +
		"удобен ли отважиться на такой выбор и выбор сей не принадлежит ли единой добродетели? Я боюсь, говорил " +
		"Плиний, лишась своего друга, я боюсь, чтобы не ослабеть в пути добродетели; уже потерял я моего проводника " +
		"и свидетеля моей жизни. Дружество по необходимости заставляет нас быть добродетельными: оно не может " +
		"храниться иначе, как между почтенными людьми, и тот, кто желает сохранить дружество, должен стараться " +
		"подражать в добродетели своему другу. Итак, в дружбе находятся и полезные советы, возбуждения последовать " +
		"добрым примерам; друг с нами делит наши печали, помогает нам в наших нуждах, не побужденный ни просьбою, " +
		"ни корыстию, ни исканием. Рассмотрим же теперь истинное свойство дружества, по коему бы можно было " +
		"познавать оное. Первое достоинство, которое должно сыскивать в друге, есть добродетель: она-то уверяет " +
		"нас в нем, что он способен к дружеству и оного достоин. Не надейся нимало на ваши обязательства, как скоро " +
		"не на сем основании они утверждены: ныне не выбор, но нужды соединяют людей, и для того-то нынешнее " +
		"дружество так же ско", "ро кончится, как и начинается: дружатся без разбору и ссорятся не раздумывая; " +
		"ничто столь не презренно: худой выбор оказывает или дурное сердце, или дурной разум. Из тысячи умей выбрать " +
		"себе друга, ничто столь не важно, как сей выбор, ибо от него зависит наше благополучие. Помысли, что нас " +
		"судят в наших друзьях; выбрать друга - значит дать обществу свой образ и открыть свое сердце. Всякий бы " +
		"трепетал, если бы размыслил, сколь важно признать кого своим другом. Хочешь ли быть почтен? Обращайся с " +
		"людьми почтенными, итак, надобно хорошо познать прежде, нежели войти в обязательствы. Первый знак, могущий " +
		"уверить, что выбранный достоин дружества, есть добродетель. Потом должно, чтобы выбираемый друг был " +
		"свободен и неподверженный страстям: одержимые гордостию и надмением не способны к таким чувствованиям, а " +
		"еще менее те, кои окованы узами любви. Любовь похищает всю цену дружества; она есть страсть жестокая, а " +
		"дружба - ощущение приятное и спокойное. Любовь напояет душу некоим восторгом радости, за коею нередко " +
		"следует жестокое сокрушение; но дружба есть наслаждение рассудка всегда чистое и всегда ровное; ничто " +
		"не может его переменить, ни ослабить; оно питает душу, и время не сокрушает, но подкрепляет его. " +
		"Занимающиеся исканием почести также не способны к сему чувствованию: они более ищут богатства, нежели " +
		"дружества, в дружестве не менее нужны чистые нравы; тот подвергается великой опасности, кто вступает в " +
		"дружбу с человеком развращенного поведения. Я думаю также, что чрезмерная молодость не способна к " +
		"наслаждению совершенным дружеством: мы видим много молодых людей, которые называются друзьями и думают, " +
		"что это истина, но их связывают одни забавы; а забавы не суть узы, достойные дружества. Мы в летах, " +
		"приличных дружеству, говорит Сенека своему другу: сильнейшие страсти в тебе угасли, остались одни " +
		"приятные, и мы будем наслаждаться забавами дружества. Но рассмотрим должности дружбы. Три время находятся " +
		"в дружбе: начало, продолжение и конец. Как всякое начало дружества наполнено чувствованиями, то все " +
		"наполнено забавою при рождающемся дружестве. Но часто случается, что вкус притупляется привычкою, " +
		"очарование исчезает, и тогда уже должно поддерживать дружество по рассуждению; подпора всегда слабая. " +
		"В дружбе, как в любви, должно щадить чувства и вкус от притупления: сия предосторожность позволена. " +
		"Но кто может воздержаться от утешения невинного и позволенного? Итак, должно утверждать дружество на " +
		"крепчайшем основании; почтение, подкрепляемое знанием достоинств; вот основание, которое никогда не " +
		"может поколебаться. Любовь пишется с завязанными глазами; но у дружества снимается сия завязка; " +
		"дружество рассматривает прежде, нежели избирает, и привязывается к одним достоинствам особы: ибо те " +
		"одни достойны быть любимы, которые в себе самих заключают причину, побуждающую их любить. Сделав добрый " +
		"выбор, надобно на нем утвердиться и иметь к своему другу непременчивое почтение и основанное не на одних " +
		"чувствах, но на рассуждении, ибо чувства могут охладеть, и тогда почтение, сопряженное со справедливостию, " +
		"только остается непременным. Не надобно дозволять себе ценить слабости в своем друге и еще менее говорить " +
		"о них; должно почитать дружбу. Но как она дана нам к тому, чтоб быть помошию добродетели, а не участвовать " +
		"в пороках, то должно уведомлять, когда друг твой уклонится с пути добродетели; если он упрямствует, " +
		"вооружись властию и силою, которую подают премудрые советы и чистота добрых намерений. Не думай, что " +
		"по окончании дружбы ты уже ничем не обязан: обязательства труднейшие остаются тебе исполнить, и в коих " +
		"одна честь может тебя поддержать; к старому дружеству должно хранить почтение. Некоторые думают, что " +
		"смерть разрушает все обязательства, и малое число людей умеют быть друзьями умерших, хотя великолепнейшее " +
		"украшение на наших похоронах суть слезы наших друзей и приятнейшая нам гробница - их сердца; не думай, " +
		"однакож, чтоб слезы, проливаемые только на погребении твоего друга, были последним знаком твоей к нему " +
		"горячности: ты еще остаешься обязанным их имени, славе и дому, он должен жить в твоем сердце, в твоем " +
		"воспоминании, в твоих ус", "тах похвалами и в твоих поступках подражанием его добродетелям. Вот правила, " +
		"вот образ дружества, начертанный на моем сердце самою природою, и я тогда только почту себя несчастливым, " +
		"когда сердце мое охладеет к толь благородным чувствованиям; когда дружество меня оставит и свет претворится " +
		"для моего сердца в пустыню."}

	for i, snt := range split(testStr, 4096) {
		if snt != testArr[i] {
			os.Exit(42)
		}
	}
}