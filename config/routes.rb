Rails.application.routes.draw do

  scope ":locale", locale: /fr/ do
    root to: "application#index"
    resources :articles
  end

  get '/:locale' => 'application#index'
  get '' => 'application#detect_language'
end
