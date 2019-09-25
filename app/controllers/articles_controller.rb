class ArticlesController < BaseController
    def index
        redirect_to(root_url)
        #render 'articles/index'
    end

    def show
        param = params[:id]
        param or not_found
        
        split = param.split('-')
        id = split[0]
        article = Rails.application.articles[params[:locale]][id]
        
        article or not_found
        
        if param != article.to_param
            redirect_to article_url(article)
        end

        params[:article] = article
    end
end
