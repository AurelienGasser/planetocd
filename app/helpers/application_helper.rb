module ApplicationHelper
    def localize_file(filename)
        idx = filename.rindex('.')
        if idx.nil?
            return filename
        end
        suffix = "_" + I18n.locale.to_s
        return filename[0..idx-1] + suffix + filename[idx..-1]
    end
end
