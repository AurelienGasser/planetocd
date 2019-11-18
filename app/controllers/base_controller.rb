class BaseController < ActionController::Base  
    layout 'default'
    before_action :redirect_domain
    before_action :set_locale

    @@redirects = {
        "planetocd.org" => 302,
        "planetetoc.fr" => 302,
        "www.planetocd.org" => 302,
        "www.planetetoc.fr" => 302,
        "www.planetetoc.org" => 302
    }

    def redirect_domain
        status = @@redirects[request.host]
        if !status.nil?
            redirect_to "#{request.protocol}#{Constants::DOMAIN}#{request.fullpath}", :status => status
        end
    end

    def set_locale
        I18n.locale = "fr"
    end

    def not_found
        head 404
        render 'errors/not_found'
    end
end
