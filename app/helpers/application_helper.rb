module ApplicationHelper
    def article_likes_count(likes_count)
      return '' if likes_count.zero?

      "(#{likes_count})"
    end
end
