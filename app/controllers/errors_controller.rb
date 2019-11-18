class ErrorsController < BaseController
    def not_found
        render 'not_found', :status => 404
    end

    def internal_error
        render 'internal_error', :status => 500
    end
end
