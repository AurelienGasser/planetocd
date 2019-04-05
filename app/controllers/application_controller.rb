class ApplicationController < BaseController
    def index
        render 'articles/index'
    end

    def about        
    end

    def detect_language
        redirect_to('/fr')
    end
end
