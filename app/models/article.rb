require 'redcarpet'

class Article
    attr_accessor :title
    
    @source = ''
    @original_title = ''
    @title = ''
    @author = ''
    @public_notes = ''
    @html_content

    def initialize(mdocd_content)
        parts = mdocd_content.split('===')
        @md_raw = parts[1]
        lines = parts[0].split("\n")
        lines.each do |line|
            if    line.start_with?('Source: ')
                @source = line[8..-1]
            elsif line.start_with?('Original title: ')
                @original_title = line[16..-1]
            elsif line.start_with?('Title: ')
                @title = line[7..-1]
            elsif line.start_with?('Author: ')
                @author = line[8..-1]
            elsif line.start_with?('Public notes: ')
                @public_notes = line[14..-1]
            else print(line)
            end
        end
    end

    def to_s
        "Source: #{@source}\n"\
        "Original Title: #{@original_title}\n"\
        "Title: #{@title}\n"\
        "Author: #{@author}\n"\
        "Public Notes: #{@public_notes}\n"\
        "Raw MD: #{@md_raw.length} characters\n" +
        render()
    end

    def html_content
        if @html_content == nil
            @html_content = render()
        end
        @html_content
    end

    def render
        markdown = Redcarpet::Markdown.new(CustomRender, fenced_code_blocks: true)
        markdown.render(@md_raw).html_safe
    end
end

class CustomRender < Redcarpet::Render::HTML
    def link(link, title, content)
        "<a href=\"#{link}\" class=\"myclass\">#{content}</a>"
    end
end
