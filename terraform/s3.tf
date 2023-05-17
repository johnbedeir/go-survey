resource "aws_s3_bucket" "s3" {
  bucket = "tfstate-comforte-prod"

  tags = {
    Environment = "production"
  }

  lifecycle {
    prevent_destroy = true
  }
}

resource "aws_s3_bucket_ownership_controls" "s3_owner" {
  bucket = aws_s3_bucket.s3.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }

  lifecycle {
    prevent_destroy = true
  }
}

resource "aws_s3_bucket_acl" "s3_acl" {
  depends_on = [aws_s3_bucket_ownership_controls.s3_owner]

  bucket = aws_s3_bucket.s3.id
  acl    = "private"

  lifecycle {
    prevent_destroy = true
  }
}

