require 'redcarpet'
require_relative '../helpers/custom_render'

class Article
    attr_accessor :id, :language, :source, :original_title, :title, :author, :public_notes_html, :content_html, :content_html_short
    
    @@markdown = Redcarpet::Markdown.new(CustomRender, fenced_code_blocks: true)
    @id = ''
    @language = ''
    @source = ''
    @original_title = ''
    @title = ''
    @author = ''
    @public_notes_md = ''
    @public_notes_html = ''
    @content_html
    @content_html_short
    @content_md = ''

    def initialize(article_path)
        mdocd_content = IO.read(article_path)
        populate_from_mocd(mdocd_content)
        convert_to_html()
    end

    def populate_from_mocd(mdocd_content)
        parts = mdocd_content.split('===')
        @content_md = parts[1]
        lines = parts[0].split("\n")
        lines.each do |line|
            if line.start_with?('Id: ')
                @id = line[4..-1]
            elsif line.start_with?('Language: ')
                @language = line[10..-1]
            elsif line.start_with?('Source: ')
                @source = line[8..-1]
            elsif line.start_with?('Original title: ')
                @original_title = line[16..-1]
            elsif line.start_with?('Title: ')
                @title = line[7..-1]
            elsif line.start_with?('Author: ')
                @author = line[8..-1]
            elsif line.start_with?('Public notes: ')
                @public_notes_md = line[14..-1]
            end
        end
    end

    def convert_to_html
        content_md_short = @content_md.split("\n\n")[0..4].join("\n\n")
        @content_html = @@markdown.render(@content_md).html_safe
        @content_html_short = @@markdown.render(content_md_short).html_safe
        if (@public_notes_md)
            @public_notes_html = @@markdown.render(@public_notes_md).html_safe
        end
    end

    def to_s
        "Id: #{@id}\n"\
        "Language: #{@language}\n"\
        "Source: #{@source}\n"\
        "Original Title: #{@original_title}\n"\
        "Title: #{@title}\n"\
        "Author: #{@author}\n"\
        "Public Notes: #{@public_notes_md}\n"\
        "Raw MD: #{@md_raw.length} characters\n"\
        "HTML: #{@content_html.length} characters\n"\
        "HTML (short): #{@content_html_short.length} characters\n"
    end
end
