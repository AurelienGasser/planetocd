class CustomRender < Redcarpet::Render::HTML
    def link(link, title, content)
        "<a href=\"#{link}\" class=\"myclass\">#{content}</a>"
    end
end
