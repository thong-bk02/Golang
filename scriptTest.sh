#!/bin/bash

echo "=== TOOL GIT AN TOAN (CO REVIEW FILE) ==="

# 1. Hiển thị trạng thái & Add file
echo "📂 Dang quet thay doi..."
git add .
echo "✅ Da Add toan bo file vao Staging."

# 2. Xử lý Commit
if [ -z "$1" ]
then
    echo "--------------------------------"
    echo "⚠️  Ban chua nhap ghi chu commit!"
    read -p "👉 Nhap ghi chu Commit: " msg
    if [ -z "$msg" ]; then
        msg="Auto update code"
    fi
else
    msg="$1"
fi

git commit -m "$msg"
echo "✅ Da Commit xong."

echo "--------------------------------"

# 3. Chọn nhánh để Push
current_branch=$(git branch --show-current)
echo "ℹ️  Ban dang dung o nhanh: $current_branch"

read -p "👉 Ban muon Push len nhanh nao? (Mac dinh: $current_branch): " target_branch

# Nếu người dùng ấn Enter luôn thì lấy nhánh hiện tại
if [ -z "$target_branch" ]; then
    target_branch=$current_branch
fi

echo "--------------------------------"

# 4. [TÍNH NĂNG MỚI] Hiển thị danh sách file sẽ đẩy đi
echo "🔍 REVIEW: Danh sach cac file se duoc day len '$target_branch':"
echo ""

# Kiểm tra xem nhánh này đã có trên server chưa
git rev-parse --verify origin/$target_branch >/dev/null 2>&1

if [ $? -eq 0 ]; then
    # Nếu nhánh đã tồn tại: So sánh Local vs Server
    # --stat: Hiển thị thống kê file | --name-status: Hiển thị tên file và trạng thái (A, M, D)
    git diff --stat origin/$target_branch..HEAD
else
    # Nếu nhánh chưa tồn tại (Push lần đầu): Hiển thị tất cả file trong commit gần nhất
    echo "⚠️  (Nhanh '$target_branch' chua co tren server. Day la lan Push dau tien)"
    echo "📄 Cac file trong commit vua roi:"
    git show --stat --oneline --name-only HEAD
fi

echo ""
echo "--------------------------------"

# 5. Hỏi xác nhận lần cuối
read -p "❓ Ban co CHAC CHAN muon push cac file tren khong? (y/n): " confirm

if [ "$confirm" == "y" ] || [ "$confirm" == "Y" ]; then
    echo "🚀 Dang Push code len '$target_branch'..."
    git push origin $target_branch
    
    if [ $? -eq 0 ]; then
        echo "🎉 THANH CONG! Code da len server."
    else
        echo "❌ THAT BAI! Kiem tra lai mang hoac quyen truy cap."
    fi
else
    echo "⛔ Da huy bo Push."
fi

echo "=== KET THUC ==="