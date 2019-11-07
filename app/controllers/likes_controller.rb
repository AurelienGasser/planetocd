class LikesController < BaseController
    def create
      Like.create(like_params)
      @likes_count = article.likes_count
    end

    private

    def like_params
      params
        .permit(:article_id)
        .merge(user_agent: request.user_agent,
               ip_address: request.remote_ip)
    end

    def article
      Rails.application.articles[I18n.locale.to_s][params[:article_id].to_i]
    end
end