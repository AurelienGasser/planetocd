class BaseController < ActionController::Base  
    layout 'default'
    before_action :set_locale

    def default_url_options
        { locale: I18n.locale }
    end

    def set_locale
        begin
            I18n.locale = params[:locale] || I18n.default_locale
        rescue
            return not_found
        end
    end

    def not_found
        raise ActionController::RoutingError.new('Not Found')
    end
end
