class BaseController < ActionController::Base  
    layout 'default'
    before_action :redirect_domain
    before_action :set_locale

    def redirect_domain
        if request.host == "planetocd.org" || request.host == "www.planetocd.org"
            redirect_to "#{request.protocol}#{Constants::DOMAIN}#{request.fullpath}", :status => :moved_temporarily
        end
    end

    def set_locale
        I18n.locale = "fr"
    end

    def not_found
        render 'errors/not_found'
    end
end
