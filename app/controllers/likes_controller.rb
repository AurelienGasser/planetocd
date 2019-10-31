class LikesController < BaseController
    def create
      Like.create(like_params)
    end

    private

    def like_params
      params
        .permit(:article_id)
        .merge(user_agent: request.user_agent,
               ip_address: request.remote_ip)
    end
end
