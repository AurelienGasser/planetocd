class ArticlesController < BaseController
    def index
    end

    def show
        article_id = params[:id]
        # TODO 404
        if article_id != 1
        end
        article = Rails.application.articles[params[:locale]][article_id]
        # TODO pass article the right way
        params[:article] = article
    end
end
