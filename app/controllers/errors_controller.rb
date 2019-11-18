class ErrorsController < BaseController
    def not_found
        render 'not_found', :status => 400
    end

    def internal_error
        render 'internal_error', :status => 500
    end
end
