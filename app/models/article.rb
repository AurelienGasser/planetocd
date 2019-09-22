require 'redcarpet'
require_relative '../helpers/custom_render'

class Article
    attr_accessor :title, :html_content, :html_content_short, :language
    
    @@markdown = Redcarpet::Markdown.new(CustomRender, fenced_code_blocks: true)
    @source = ''
    @id = ''
    @language = ''
    @original_title = ''
    @title = ''
    @author = ''
    @public_notes = ''
    @html_content
    @html_content_short

    def initialize(article_path)
        mdocd_content = IO.read(article_path)
        populate_from_mocd(mdocd_content)
        convert_to_html()
    end

    def populate_from_mocd(mdocd_content)
        parts = mdocd_content.split('===')
        @md_raw = parts[1]
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
                @public_notes = line[14..-1]
            end
        end
    end

    def convert_to_html
        @html_content = @@markdown.render(@md_raw).html_safe
        md_raw_short = @md_raw.split("\n\n")[0..4].join("\n\n")
        @html_content_short = @@markdown.render(md_raw_short).html_safe
    end

    def to_s
        "Id: #{@id}\n"\
        "Language: #{@language}\n"\
        "Source: #{@source}\n"\
        "Original Title: #{@original_title}\n"\
        "Title: #{@title}\n"\
        "Author: #{@author}\n"\
        "Public Notes: #{@public_notes}\n"\
        "Raw MD: #{@md_raw.length} characters\n"\
        "HTML: #{@html_content.length} characters\n"\
        "HTML (short): #{@html_content_short.length} characters\n"
    end
end
