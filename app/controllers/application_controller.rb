class ApplicationController < BaseController
    def index
    end

    def detect_language
        redirect_to('/fr')
    end
end
