Rails.application.routes.draw do

  # scope ":locale", locale: /fr/ do
    root to: "application#index"
    resources :articles
    get 'about', to: 'static_pages#about', as: :about 
  # end

  get '/404', to: "errors#not_found"
  get '/500', to: "errors#internal_error"
end
