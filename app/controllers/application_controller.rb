class ApplicationController < BaseController
    def index
        redirect_to(articles_url)
        #render 'articles/index'
    end

    def about        
    end

    def detect_language
        redirect_to('/fr')
    end
end
