Rails.application.routes.draw do

  scope ":locale", locale: /fr/ do
    root to: "application#index"
    resources :articles
    get 'about', to: 'static_pages#about', as: :static 
  end

  get '/:locale' => 'application#index'
  get '' => 'application#detect_language'
end
