{% macro generatewhere(where, n) %}
    sqls{{n}} := []string{}
    {% if where.Test != "true" %}
    if {{where.Test }}{
    {% endif %}        
        {% for item in where.Itemlist %}
            {% if item|isallatom %}
                sqls{{n}} = append(sqls{{n}}, `{{item|generateatom}}`)
            {% elif item.Operate != "atom" %}
            {
               {{generatewhere(item, n+1)}}
               sqls{{n}} = append(sqls{{n}}, strings.Joins(sqls{{n+1}}, " {{where.Operate}} "))
            }
            {% elif  item.Test == "true"%}
              sqls{{n}} = append(sqls{{n}}, `({{item.Sql}})`)
            {% else %}
                if {{item.Test}} {
                    sqls{{n}} = append(sqls{{n}}, `({{item.Sql}})`)
                 }
            {% endif %}
        {% endfor %}        
    {% if where.Test != "true" %}
    }
    {% endif %}

    {% if n==0 %}
             where := strings.Joins(sqls0, " {{where.Operate}} "))
    {% endif %}    
{% endmacro %}

func funcname() {
    {% if where|isallatom %}
        where = `{{where|generateatom}}`
    {% else %}
        {{generatewhere(where, 0)}}
    {% endif %}
}
