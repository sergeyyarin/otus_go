package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = false

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var text2 = `Overall logic is simple and this process are the same (or at least very similar) in every technology.
i) A UE is in connection with a cell (let's call this 'Cell A').
ii) Now a situation that requires handover happened.
iii) Network send "signal quality measurement" command to UE for the garget cell ('Cell B') to which it will handover to.
iv) UE performance the measurement and report the "measurement result" to the network via the current cell (Cell A).
v) Network evaluate the measurement result reported by UE.
vi) If the evaluation result turns out to be good for handover, Network send 'Change Cell' command to UE.
vii) UE perform the cell change process.
viii) If cell change process is completed properly, UE send 'cell change completion' message to the network via the 
target cell (Cell B).

I used very generic term e.g, "signal quality measurement command", "measurement result", "Change Cell Command", 
"Cell Change Completion Message" etc. These generic commands can be translated to a specific jargon for each technology.
 For example, if I translate these for UMTS, they would be as follows :
"Signal quality measurement command" ==> Measurement Control
"Measurement Result" ==> Measurement Report
"Change Cell Command" ==> Physical Channel Reconfiguration or ActiveSetUpdate
"Cell Change Compeletion Message ==> Physical Channel Reconfiguration Complete or ActiveSetUpdateComplete

If you translate them into LTE jargon, they will be as follows.
"Signal quality measurement command" ==> RRC Connection Reconfiguration
"Measurement Result" ==> Measurement Report
"Change Cell Command" ==> RRC Connection Reconfiguration
"Cell Change Compeletion Message ==> RRC Connection Reconfiguration Complete

You may noticed that LTE is using the same message called "RRC Connection Reconfiguration" both for "Signal quality 
measurement command" and "Change Cell Command". How UE can tell whether it means "Signal quality measurement command"
 or "Change Cell Command" ?
Good question ! You will see the answer later.

Then you may have whole lots of questions. It is very good. The more questions you have, the more information 
you will get through this page.. (not now, in the future -: ) Following is a set of my personal questions.
i) you talked about "Signal Quality Measurement". What kind of signal quality UE has to measure ?
 Would it be a certain absolute value ? or a some relative value with reference to some other value ?
  or is it a special event changes ?
ii) How much time I can leave the current cell to perform the measurement for target cell ?
 (If the leave too long from the current cell to measure target cell, the call would drop. 
	But if this time is too short, UE would not get correct measurement values).
iii) What if UE failed to performe the measurement or fail to find the target cell ?
iv) you talked about "Change Cell", how UE can change cell ? Just cut the connection with the current cell
 and reconnect to the target cell ? or is there any specific procedure ?
v) Cutting the connection from the current cell will be easy, but how can UE reconnect to target cell ?
vi) What if UE failed to reconnect to target cell after he cut off the connection with the current cell ?`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})

	t.Run("extra test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			//
		} else {
			expected := []string{
				"the",
				"cell",
				"to",
				"?",
				"UE",
				"measurement",
				"is",
				"==>",
				"a",
				"or",
			}
			require.Equal(t, expected, Top10(text2))
		}
	})
}
