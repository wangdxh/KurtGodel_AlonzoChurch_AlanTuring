 
 
{% macro generatelist(structef, n, ismregelist)%}
    

{% endmacro %}

// from type get structdef.StructItem
// getitemtag unique assign
// unique left, right

func Groupby({{mergefrom.Name}} []{{mergefrom.Type}}) (merge {% if mergeto.Islist %}[]{% endif %} {{mergeto.Type}}) {
	for _, valres := range {{mergefrom.Name}} {

		var m0 *{{mergeto.Type}}
		for inx := 0; inx < len(merge); inx++ {
            // unique left, right
			if merge[inx].TestA.Id == valres.TestA.Id {
				m0 = &merge[inx]
				break
			}
		}
		if m0 == nil {
			merge = append(merge, TestMerge{})
			m0 = &merge[len(merge)-1]

            {% for item in mergeto.Type|structdefgetStructItem %}
                {% if item.Islist == false %}
                    m0.{{Item.Name}} = valres.{{item|tagassign}}
                {% endif %}
            {% endfor %}
			m0.TestA = valres.TestA
		}

        {% for item in mergeto.Type|structdefgetStructItem%}
            {% if item.Islist == true %}
                m0.{{Item.Name}} = valres.{{item|tagassign}}
            {% endif %}
        {% endfor %}



		var m1 *TestBMerge
		for inx := 0; inx < len(m0.B); inx++ {
			if m0.B[inx].TestB.Id == valres.TestB.Id {
				m1 = &m0.B[inx]
				break
			}
		}

		if m1 == nil {
			m0.B = append(m0.B, TestBMerge{})
			m1 = &m0.B[len(m0.B)-1]

			m1.TestB = valres.TestB
		}

		var m2 *TestC
		for inx := 0; inx < len(m1.C); inx++ {
			if valres.TestC.Id == m1.C[inx].Id {
				m2 = &m1.C[inx]
				break
			}
		}
		if m2 == nil {
			m1.C = append(m1.C, TestC{})
			m2 = &m1.C[len(m1.C)-1]

			*m2 = valres.TestC
		}

	}
	return
}
