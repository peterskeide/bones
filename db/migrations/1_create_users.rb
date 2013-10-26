Sequel.migration do
  change do
    create_table(:users) do
      primary_key :id
      String      :password, null: false, size: 255
      String      :email,    null: false, size: 255
    end
  end
end
